#!/usr/bin/env node

import parseArgs from "minimist"
import njk from "nunjucks"
const { Environment, FileSystemLoader } = njk
const { render } = njk
import { $ } from "zx"

const argsConfig = { default: { f: "pdf" } }
const args = parseArgs(process.argv.slice(2), argsConfig)
const source = args._[0].replace(/\.lys$/, "")
const tags = { commentStart: "<#", commentEnd: "#>" }
const env = new Environment(new FileSystemLoader("."), { tags })
const score = env.render(`${source}.lys`)
const lilypond = $`lilypond -f ${args.f} -o ${source} -`
lilypond.stdin.write(score)
lilypond.stdin.end()
