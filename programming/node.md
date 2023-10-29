# Node.js

## Reactor pattern

- **Sync Event Demultiplexer** (SED) = provides an OS-level syscall for async,
  non-blocking IO operations e. g. `epoll` on Linux. SED watches sync for a set
  of async operations to complete and allows multiple async operations to be
  processed in a single thread (events demultiplexing)
- **Reactor pattern** = async executes a handler (app callback) for each async
  operation (non-blocking IO)
    - App requests async operations [resource (file), operation (read), handler
      (callback)] at SED (non-blocking)
    - SED watches requested async operations for completion (blocking)
    - SED enqueues [event (operation), handler (callback)] to the event queue
      (EQ)
    - Single-threaded event loop (EL) reads the EQ and executes handlers (app
      callbacks) to completion (no race conditions). App callbacks request more
      async operations at SED
    - When a callback requests a new async operations it gives control back to
      the event loop. Invoking an async operation always unwinds the stack back
      to the event loop, leaving it free to handle other requests (IO-bound)
    - EL blocks again (next EL cycle) at SED for new async operations to
      complete
- **libuv** = SED + reactor (event loop + event queue) = cross-platform
  low-level IO engine for Node.js
    - High resolution clock and timers
    - Async thread pool and communication channels
    - Async child processes and signal handling
    - Async FS, TCP/UDP, DNS
    - Async IPC via shared UNIX domain sockets (bidirectional channels)
- **Node.js** (async IO operations in parallel by libuv threads, sync callback
  code to completion in the event loop) = libuv (SED + reactor) + V8 (JavaScript
  runtime) + modules
- In Node.js only async, non-blocking IO operations are **executed in parallel**
  by the **libuv internal IO threads**. Sync code inside callbacks (until next
  async operation) is **executed concurrently to completion without race
  conditions** by the **single-threaded event loop**
- Node.js scales well (high throughput) when serving **large number of clients**
  with a **small number of threads**: the main thread running the event loop,
  libuv aync IO threads, CPU-intensive Worker threads. Node.js scales well when
  processing **small async IO tasks** and **small CPU-intensive tasks**.
- **Partition big tasks** into smaller sub-tasks to be processed concurrently
  and **minimize variation** of tasks to give fair amount of time to every
  client. Alternatively offload CPU-intensive tasks to **Worker threads** paying
  communication **cost of serialization and de-serialization**. IO threads are
  blocked waiting for async IO operations to complete. Worker threads are
  executed in parallel on CPU cores
    ```js
    // large blocking task
    function sumFirstN(n) {
      let sum = 0
      for (let i = 1; i <= n; ++i) { sum += i }
      return sum
    }
    // partitioned small tasks (callback)
    function sumFirstN2(n, { size = 2 }, done) {
      let sum = 0, i = 1
      function next(m = size) {
        while (i <= n && m-- > 0) { sum += i++ }
        i > n ? done(null, sum) : queueMicrotask(next)
      }
      queueMicrotask(next) // must be async
    }
    // partitioned small tasks (Promise)
    function sumFirstN3(n, { size = 2 } = { }) {
      return new Promise(resolve => {
        let sum = 0, i = 1
        function next(m = size) {
          while (i <= n && m-- > 0) { sum += i++ }
          i > n ? resolve(sum): queueMicrotask(next)
        }
        next() // Promise is always async
      })
    }
    console.log(sumFirstN(10)) // 55
    sumFirstN2(10, { size: 3 }, (err, sum) => console.log(sum)) // 55
    console.log(await sumFirstN3(10, { size: 3 })) // 55
    ```

## Callback pattern

- **Callback** = **first-class function** (to pass a callback) + **closure** (to
  retain the caller context) that is passed to an async function (which
  **returns immediately**, non-blocking) and **is executed on completion** of
  the requested async operation. A callback implements the Continuation-Passing
  Style, CPS = control is passed explicitly in a form of a
  continuation/callback, `return/throw` => `callback(error, result)`
- Callback **is expected to be invoked exactly once** either with a result or an
  error. A callback has **trust issues**: a caller passes callback code to a
  third party, an async function receives callback code from a third party
- Async result/error is **propagated through a chain** of nested
  `callbacks(error | null, data | cb(data))` that accept an `error | null`
  **first**, then a `result | callback(data)` that comes always **last**. Never
  `return` a result or `throw` an error from a callback (`return` and `throw`
  are unusable from a callback)
- Sync CPU-bound operations in a callback are discouraged as they block the
  event loop and put on hold all concurrent processing blocking the whole
  application
    - Interleave each step of a CPU-bound sync operation with `queueMicrotask()`
      or `setImmediate()` to let other IO tasks to be processed by the event
      loop (not efficient due to context switching, more complex interleaving
      algorithm)
    - Pool of reusable external Node.js processes `child_process.fork()` with a
      communication channel `process/worker.send().on(message)` leaves the event
      loop unblocked (efficient: reusable processes run to completion,
      unmodified algorithm)
    - `new Worker(module, opts)` see below
- Avoid mixing sync with **async callback behavior** under the same interface.
  To convert sync callback to an async callback use
    - **Next tick queue** `process.nextTick()` (managed by Node.js,
      non-clearable) is executed just after the current operation, before other
      IO tasks in **the same cycle** of the event loop
    - **Mictrotask queue** `queueMicrotask()` (managed by V8, favor over
      `process.nextTick()`) is used to execute Promise handlers `.then()`,
      `.catch()`, and `.finally()`. The microtask queue is drained immediately
      after the next tick queue is drained within **the same cycle** of the
      event loop
        ```js
        import { EventEmitter } from "node:events"
        class EE extends EventEmitter {
          constructor() {
            super()
            // emits sync before a listener is attached (does not work)
            this.emit("ready")
            // amits async after a listener is attached (works)
            queueMicrotask(() => this.emit("ready"))
          }
        }
        const ee = new EE()
        ee.on("ready", () => console.log("ready"))
        ```
    - **Macrotask queue** is always **the next cycle** of the event loop
        - **Immediate task** `setImmediate()` is executed as soon as possible on
          the next cycle of the event loop, before any timers `setTimeout()` and
          `setInterval()`
        - **Delayed task** `setTimeout()` is executed after a delay in one of
          the next cycles of the event loop
- An uncaught error thrown from an async function propagates to the stack of the
  event loop (not to the next callback, not to the stack of the caller that
  triggered the async operation), `process.on(unchaughtException)` is
  emitted and the process exits with a non-zero exit code
- **Callback hell** = deeply nested code as a result of **in-place nested
  anonymous callbacks** that unnecessary consume memory because of closures
    - Do not abuse in-place nested anonymous callbacks
    - **Create named callbacks** with clearly defined interface (and without
      unnecessary closures that consume memory)
- **Sequential execution** (CPS) = executes async tasks in sequence one
  task at a time until all complete or the first error
    ```js
    function task(v, done) {
      setTimeout(
        () => v ? (console.log(v), done(null, v)) : done(new Error("oh")), 500
      )
    }
    function sequence(tasks, done) {
      let i = 0
      function next(err) {
        if (err) { return done(err) }
        i < tasks.length ? tasks[i++](next) : done(null)
      }
      queueMicrotask(next)
    }
    const tasks = [1, 2, 3].map(el => done => task(el, done)) // 1, 2, 3, null
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done)) // 1, 2, oh
    sequence(tasks, console.log)
    sequence([], console.log) // null
    ```
- **Parallel execution** (sync loop through all tasks) = executes tasks in
  parallel (unlimited) until all complete or the first error
    ```js
    function parallel(tasks, done) {
      let completed = 0
      function next(err) {
        if (err) { return done(err) }
        if (++completed === tasks.length) { return done(null) }
      }
      if (tasks.length === 0) { queueMicrotask(() => done(null)) }
      for (const task of tasks) { task(next) }
    }
    const tasks = [1, 2, 3].map(el => done => task(el, done)) // 1, 2, 3, null
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done)) // 1, 2, oh, 3
    parallel(tasks, console.log)
    parallel([], console.log) // null
    ```
- **Limited parallel execution** (sync loop until a limit) = executes at most N
  tasks in parallel until all complete or the first error
    ```js
    function parallelLimit(tasks, { limit = 2 }, done) {
      let completed = 0, running = 0, i = 0
      function next(err) {
        if (err) { return done(err) }
        if (++completed === tasks.length) { return done(null) }
        if (--running < limit) { parallel() }
      }
      function parallel() {
        while (i < tasks.length && running < limit) {
          tasks[i++](next); ++running
        }
      }
      if (tasks.length === 0) { queueMicrotask(() => done(null)) }
      parallel()
    }
    // 1, 2 | 3, 4 | 5, null
    const tasks = [1, 2, 3, 4, 5].map(el => done => task(el, done))
    // 1, 2 | 3, oh | 4
    const tasks = [1, 2, 3, 0, 4].map(el => done => task(el, done))
    parallelLimit(tasks, { limit: 2 }, console.log)
    parallelLimit([], { }, console.log) // null
    ```

## Observer pattern

- **Observer** = `EventEmitter` continuously notifies **multiple
  observers/listeners** on its state changes through different types of
  recurrent events. `EventEmitter` **does not have trust issues** = a callback
  is controlled by a caller
- Events are **guaranteed to fire async** in the next cycle of the event loop.
  However, `ee.emit()` is called sync for all registered listeners. Combining
  `EventEmitter` with the callback interface is an elegant and flexible solution
- `EventEmitter` = registers multiple `function listener(event, ...args)` for
  specific event types `ee.on|once(event, listener)`. Synchronously
  `ee.emit(event, ...args)` for all registered listeners.
  `ee.removeListener(event, listener)` unsubscribes a listeners when it is no
  longer needed to avoid memory leaks due to captured context in listener
  closures. If `ee.on(error, ...args)` is not registered an `Error` is thrown.
  Always register a listener for the `error` event
- Async result/error is propagated through `emit` events. Never `return` a
  result or `throw` an error from an `EventEmitter` (`return` and `throw` are
  unusable with `EventEmitter`)
    ```js
    import { EventEmitter } from "node:events"
    class EE extends EventEmitter { }
    const ee = new EE()
    ee.on("error", err => console.error(err.message))
    ee.on("start", console.log)
    ee.on("start", function () { this.emit("error", new Error("oh")) })
    setTimeout(() => ee.emit("start", 1), 500) // 1, oh
    ```

## Promise pattern

- Sync code has an **ever-growing** list of `.on(uncaughtExeption)`. Promise
  code has a **growing-and-shrinking** list of `.on(unhandledRegection)` as a
  rejection can be handled later when a rejected promise gets a rejection
  handler
- **Sequential iteration** (dynamic promise chain in a sync loop)
    ```js
    function task(v) {
      return new Promise((resolve, reject) =>
        setTimeout(
          () => v ? (console.log(v), resolve()) : reject(new Error("oh")), 500
        )
      )
    }
    function sequence(tasks) {
      let p = Promise.resolve()
      for (const task of tasks) { p = p.then(task) }
      return p
    }
    // 1, 2, 3, undefined
    const tasks = [1, 2, 3].map(el => () => task(el))
    // 1, 2, oh
    const tasks = [1, 2, 0, 3].map(el => () => task(el))
    sequence(tasks).then(console.log, console.error)
    sequence([]).then(console.log, console.error) // undefined
    ```
- **Parallel execution** (sync loop for all tasks)
    ```js
    function parallel(tasks) {
      return new Promise ((resolve, reject) => {
        let completed = 0
        function next() {
          if (++completed === tasks.length) { return resolve() }
        }
        if (tasks.length === 0) { return resolve() }
        for (const task of tasks) { task().then(next, reject) }
      })
    }
    // 1, 2, 3, undefined
    const tasks = [1, 2, 3].map(el => () => task(el))
    // 1, 2, oh, 3
    const tasks = [1, 2, 0, 3].map(el => () => task(el))
    parallel(tasks).then(console.log, console.error)
    parallel([]).then(console.log, console.error) // undefined
    ```
- **Limited parallel execution** (sync loop until a limit)
    ```js
    function parallelLimit(tasks, limit = 2) {
      return new Promise((resolve, reject) => {
        let completed = 0, running = 0, i = 0
        function next() {
          if (++completed === tasks.length) { return resolve() }
          if (--running < limit) { parallel() }
        }
        function parallel() {
          while (i < tasks.length && running < limit) {
            tasks[i++]().then(next, reject); ++running
          }
        }
        if (tasks.length === 0) { return resolve() }
        parallel()
      })
    }
    // 1, 2 | 3, 4 | 5, undefined
    const tasks = [1, 2, 3, 4, 5].map(el => () => task(el))
    // 1, 2, | 3, oh | 4
    const tasks = [1, 2, 3, 0, 4].map(el => () => task(el))
    parallelLimit(tasks, 2).then(console.log, console.error)
    parallelLimit([]).then(console.log, console.error) // undefined
    ```

## Async/await pattern

- `async` function **always returns a Promise** immediately and synchronously.
  `try/catch/throw` inside an `async` function works for both sync and async
  code. Use `return await` to prevent errors on the caller side and `catch`
  errors locally
    ```js
    async function localError() {
      try { return task(false) } // Caller oh
      try { return await task(false) } // Local oh
      catch (e) { console.error("Local", e.message) }
    }
    localError().catch(e => console.error("Caller", e.message))
    ```
- `async` function is a `Promise`-`yield`ing generator that on each `await`
  expression `yield`s a Promise and suspends the generator, its internal state
  is saved and the control is retuned to the event loop. When the Promise that
  has been `await`ed settles, the control is given back to the `async` function
  and the generator is resumed
- `Promise` abstraction and `async/await` syntax is used to manage async
  operations in a sync-like manner. However, if async operations are unrelated,
  `await` introduces unnecessary blocking. Do not use `await` inside a loop, use
  `Promise.all()` instead
    ```js
    console.log(await task(1), await task(2)) // sequence, slow
    console.log(await Promise.all([task(1), task(2)])) // parallel, fast
    ```
- **Sequential iteration** (sync loop with await)
    ```js
    async function sequence(tasks) {
      for (const task of tasks) { await task() }
    }
    // 1, 2, 3, undefined
    const tasks = [1, 2, 3].map(el => done => task(el, done))
    // 1, 2, oh
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done))
    try {
      console.log(await sequence(tasks))
      console.log(await sequence([])) // undefined
    } catch (e) { console.error(e.message) }
    ```
- **Parallel execution** (await loop)
    ```js
    // Use the prallel() from the Promise pattern above
    // 1, 2, 3, undefined
    const tasks = [1, 2, 3].map(el => done => task(el, done))
    // 1, 2, oh, 3
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done))
    try {
      console.log(await parallel(tasks))
      console.log(await parallel([])) // undefined
    } catch (e) { console.error(e.message) }
    ```
- **Limited parallel execution**
    ```js
    // Use the parallelLimit() from the Promise pattern above
    // 1, 2 | 3, 4 | 5, undefined
    const tasks = [1, 2, 3, 4, 5].map(el => done => task(el, done))
    // 1, 2, | 3, oh | 4
    const tasks = [1, 2, 3, 0, 4].map(el => done => task(el, done))
    try {
      console.log(await parallelLimit(tasks, 2))
      console.log(await parallelLimit([], 2)) // undefined
    } catch (e) { console.error(e.message) }
    ```

## Stream pattern

- **Streaming** = parallel staged processing of data in chunks as soon as it
  arrives with internal buffering and **backpressure** (reactive, modular,
  composable, **constant memory**, short GC cycles). `Stream` is an abstraction
  on top of a data source `Readable`, a data transformation `Transform`, and a
  data sink `Writable`. `Stream` extends `EventEmitter`
    - **Binary mode** `Buffer` or string with an encoding for IO processing
    - **Object mode** JavaScript object/array for function composition
- `Writable` = standard abstraction of a **data sink** on top of an underlying
  resource with **backpressure** when an internal buffer has exceeded the
  `highWaterMark` then `.write(chunk) => false` stop writing until a Writable
  will notify when an underlying resource is ready for writing `.on(drain)` to
  resume writing
    ```js
    import { Writable } from "node:stream"
    import { finished } from "node:stream/promises"
    class Sink extends Writable {
      _construct(done) { // allocates resources
        this.buffer = []; done(null)
      }
      _write(chunk, encoding, done) { // if error done(new Error("oh"))
        setTimeout(() => { this.buffer.push(chunk); done(null) }, 100)
      }
      _final(done) { // flushes buffered data before a Writable end
        this.buffer = this.buffer.join(""); done(null)
      }
      _destroy(err, done) { // disposes resources
        this.buffer += "."; done(err)
      }
    }
    const sink = new Sink()
    sink.write("a"); sink.write("b"); sink.end("c")
    await finished(sink)
    console.log(sink.buffer) // abc.
    ```
- `Readable` = standard abstraction of a **data source** from an underlying
  resource with **backpressure** when `this.push(chunk) => false` stop reading
  from the underlying resource. The `_read(size)` will be called later to read
  more data from the underlying resource. A Readable will start reading from the
  underlying resource only when data consumption begins
    ```js
    import { Readable } from "node:stream"
    class Source extends Readable {
      constructor(source, { encoding = "utf8", ...opts } = { }) {
        super({ encoding, ...opts })
        this.source = source; this.i = 0
      }
      _construct(done) { // allocates resources
        this.source = this.source.split(" "); done(null)
      }
      _read(size) {
        setTimeout(() => // if error this.destroy(new Error("oh"))
          this.i < this.source.length ? this.push(this.source[this.i++]) :
            this.push(null), 100 // end of a Readable
        )
      }
      _destroy(err, done) { // disposes resources
        this.source = null; done(err)
      }
    }
    ```
- **Async interator** `for await ... of` to consume a Readable
    ```js
    const source = new Source("a b c")
    for await (const chunk of source) { console.log(chunk) } // a b c
    ```
- **Flowing mode** (push) = pushes data as soon as it arriaves. The flowing mode
  is activated by `.on(data)`, `.resume()`, `.pipe(writable)`
    ```js
    const source = new Source("a b c")
    source.on("data", chunk => console.log(chunk))
    source.on("end", () => console.log(".")) // a b c .
    ```
- **Paused mode** (pull) = pulls data in a controlled way. The paused mode is
  activated by `.on(readable)`, `.pause()`, `.unpipe(writable)`
    ```js
    const source = new Source("a b c")
    source.on("readable", () => {
      let chunk
      while(chunk = source.read()) { console.log(chunk) }
    })
    source.on("end", () => console.log(".")) // a b c .
    ```
- **Piping** = `readable.pipe(duplex | writable)` creates a stream chain,
  switches a Readable to the flowing mode, returns the last stream for chaining,
  multiple Writables can be attached to the same Readable
    ```js
    const source = new Source("a b c"),
          sink = new Sink(), sink2 = new Sink()
    source.pipe(sink); source.pipe(sink2)
    await finished(sink); await finished(sink2)
    console.log(sink.buffer, sink2.buffer) // abc. abc.
    ```
- `Duplex` = **independent** Readable `_read(size)` and Writable `_write(chunk,
  encoding, done)`, `_final(done)` for stream chaining through
  `readable.pipe(duplex | writable)` or `pipeline(readable, ...transform,
  writable)`
- `Transform` = a Readable **dependent** on Writable through a transformation
  that follows a pattern Writable => Transform => Readable.
  `readable.pipe(duplex | writable)` returns the last stream for stream
  chaining, controls backpressure automatically. Errors are not propagated
  automatically through a `pipe()`, `on(error)` handlers must be attached to
  every step. Destruction of a pipeline constructed with `.pipe()` has to be
  performed manually
    ```js
    import { Transform } from "node:stream"
    class Double extends Transform {
      _transform(chunk, encoding, done) { // Writable side
        setTimeout(() => { // if error done(new Error("oh"))
          const trans = String(chunk).split("").map(ch => ch + ch).join("")
          this.push(trans); done(null) // Readable side
        }, 100)
      }
      _flush(done) { this.push("-"); done(null) } // before a Readable end
    }
    class Upcase extends Transform {
      _transform(chunk, encoding, done) {
        setTimeout(() => {
          this.push(String(chunk).toUpperCase()); done(null)
        }, 100)
      }
      _flush(done) { this.push("_"); done(null) }
    }
    const source = new Source("a b c"), upcase = new Upcase(),
          double = new Double(), sink = new Sink()
    await finished(source.pipe(double).pipe(upcase).pipe(sink))
    console.log(sink.buffer) // AABBCC-_.
    ```
- `pipeline(readable, ...transform, writable)` = combines streams in a
  **non-composable** end-to-end pipeline that follows a pattern Readable =>
  Writable, automatically handles backpressure, errors, and destruction of
  streams on the pipeline success or failure
    ```js
    import { pipeline } from "node:stream/promises"
    const source = new Source("a b c"), upcase = new Upcase(),
          double = new Double(), sink = new Sink()
    await pipeline(source, double, upcase, sink)
    console.log(sink.buffer) // AABBCC-_.
    ```
- `compose(...streams)` = combines streams in a new **composable** Duplex stream
  that follows a pattern Writable => Readable using the `pipeline()`
    ```js
    import { compose } from "node:stream"
    const source = new Source("a b c"), upcase = new Upcase(),
          double = new Double(), sink = new Sink()
    await pipeline(source, compose(double, upcase), sink)
    console.log(sink.buffer) // AABBCC-_.
    ```
- **Async iterator** `for await ... of` `next() => Promise()` and
  **async generator** `async function* ...` `yield Promise()` form a basis for
  **language-level construction of streams**
    ```js
    async function* upcase(values) {
      for await (const value of values) {
        yield new Promise(resolve =>
          setTimeout(() => resolve(value.toUpperCase()), 100)
        )
      }
    }
    const source = new Source("a b c"), sink = new Sink()
    await pipeline(source, upcase, sink)
    console.log(sink.buffer) // ABC.
    ```

## Child process

- `spawn(cmd, args, opts): ChildProcess` executes a command from a PATH. Async
  foundation for all other tools
    ```js
    import { spawn } from "node:child_process"
    const ls = spawn("ls", ["-lah", "/usr/lib"])
    ls.on("error", error => console.error(error))
    ls.on("close", exitCode => console.log(exitCode))
    ls.stdout.setEncoding("utf8")
    ls.stdout.on("data", chunk => console.log(chunk))
    ls.stderr.on("data", chunk => console.log(chunk))
    ```
- `exec(cmd, opts, done): ChildProcess` executes a shell command
    ```js
    import { exec } from "node:child_process"
    exec("for x in a b c; do echo $x; done", (err, stdout, stderr) => {
      if (err) { return console.log(err, stderr) }
      console.log(stdout, stderr) // a b c
    })
    ```
    ```js
    import { promisify } from "node:util"
    import { exec } from "node:child_process"
    const execp = promisify(exec)
    try {
      const { stdout, stderr } =
            await execp("echo a b c | tr '[:lower:]' '[:upper:]'")
      console.log(stdout, stderr) // A B C
    } catch (err) { console.error(err) }
    ```
- `execFile(file, args, opts, done): ChildProcess` executes a command from a
   file without a shell
    ```js
    import { execFile } from "node:child_process"
    execFile("/usr/bin/node", ["--version"], (err, stdout, stderr) => {
      if (err) { return console.error(err, stderr) }
      console.log(stdout, stderr) // v20.8.1
    })
    ```
- `fork(module, args, opts): ChildProcess` forks a Node.js child process with a
  bidirectional IPC channel
    ```js
    import { fork } from "node:child_process"
    const [node, file, args] = process.argv
    if (args === "child") { // child
      // recieve a signal and close the IPC channel with a parent
      process.on("SIGUSR2", () => process.disconnect())
      process.on("message", msg => { // receive a message from a parent
        console.log("chd", msg)
        process.send({ res: "chd => par" }) // send a amessage to a parent
      })
    } else { // parent
      const child = fork(file, ["child"]) // fork a Node.js child process
      child.on("error", error => console.error("par", error))
      child.on("close", exitCode => console.log("par", exitCode))
      // send a message to a child once a child is spawnedk
      child.on("spawn", () => child.send({ req: "par => chd" }))
      child.on("message", msg => { // receive a message from a child
        console.log("par", msg)
        child.kill("SIGUSR2") // send a signar to a child
      })
    }
    ```

## Worker thread

- `new Worker(module, opts)` creats an independnet dedicated thread for
  CPU-bound tasks executed in paralle with the main thread event loop. A
  Worker has an async bidirectional communication chanel with its parent
  `parentPort/worker.postMessage()/.on(message)`, per-thread own event loop,
  and V8 instance (small memory footprint, fast startup time, safe: no
  syncronization, no resource sharing). All communication over a channel between
  a parent and a worker is serialized and deserializedk
    ```js
    import {
      Worker, isMainThread, parentPort, workerData,
      setEnvironmentData, getEnvironmentData
    } from "node:worker_threads"
    const [node, file] = process.argv
    if (isMainThread) { // main thread
      // parameterize a worker before creation of a worker
      setEnvironmentData("worker.inc", 3);
      // provide a workload for a worker
      const worker = new Worker(file, { workerData: [1, 2, 3] })
      worker.on("error", error => console.error("main", error))
      worker.on("exit", exitCode => console.log("main", exitCode))
      // receive a message from a worker
      worker.on("message", msg => console.log("main", msg))
      // send a message to a worker
      setTimeout(() => worker.postMessage({ req: "start" }), 100)
    } else { // worker thread
      const inc = getEnvironmentData("worker.inc") // get a clone of parameters
      console.log("work", workerData, inc)
      // receive a clone of a message from a parent
      parentPort.on("message", msg => {
        console.log("work", msg)
        // process a clone of a workload and send a message to a parent
        parentPort.postMessage(workerData.map(el => el + inc))
        parentPort.close() // close an async bidirectional channel
      })
    }
    ```
- A custom `MessageChannel` can be created on either thread for separation of
  concerns and one of the `MessagePort`s can be passed to the other thread over
  the default channel
    ```js
    import { Worker, isMainThread, parentPort, MessageChannel }
    from "node:worker_threads"
    const [node, file] = process.argv
    if (isMainThread) {
      const worker = new Worker(file)
      worker.on("error", error => console.error("main", error))
      worker.on("exit", exitCode => console.log("main", exitCode))
      // create a new dedicate channel for separation of concerns
      const newChannel = new MessageChannel()
      // send a new channel to a worker
      worker.postMessage({ port: newChannel.port1 }, [newChannel.port1])
      // receive a message from a worker on a new channel
      newChannel.port2.on("message", msg => console.log("main", msg))
    } else {
      let newPort
      parentPort.once("message", msg => {
        // receive a new channel from a parent
        ({ port: newPort } = msg)
        // send a message to a parent over a new channel
        newPort.postMessage({ worker: "hi" })
        newPort.close()
      })
    }
    ```
