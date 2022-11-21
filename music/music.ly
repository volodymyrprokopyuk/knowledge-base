\version "2.23.81"

chNameExceptions = {
  % Triads
  <c e g>-\markup \super "△"
  <c ees g>-\markup \super "−"
  <c e gis>-\markup \super "+"
  % Seventh chords
  <c e g b>-\markup \super "△7"
  <c ees g bes>-\markup \super "−7"
  <c e gis bes>-\markup \super "+7"
  <c ees ges bes>-\markup \super "ø7"
}

\layout {
  \context {
    \ChordNames
    chordNameExceptions =
    #(append (sequential-music-to-chord-exceptions chNameExceptions #t)
      ignatzekExceptions)
  }
}

#(define (replace-step repl pitches)
  (map (lambda (pitch)
    (if (eqv? (ly:pitch-steps pitch) (ly:pitch-steps repl)) repl pitch))
  pitches))

chordmodifiers.hdim =
  #(lambda (pitches) (replace-step #{ ees' #} (replace-step #{ ges' #} pitches)))

chSeq = \chordmode {
  \key c \major
  c1 c:m c:aug c:dim \break
  c1:maj7 c:m7 c:aug7 c:dim7 c:7 c:m7.5- c:hdim7
}

<<
  \new Staff = cMajor \chSeq
  \new ChordNames = cMajor \chSeq
>>

cMajor = \chordmode {
  \key c \major
  c1:maj7 d:m7 e:m7 f:maj7 g:7 a:m7 b:hdim7 \bar "|."
}

gMajor = \chordmode {
  \key g \major
  g1:maj7 a:m7 b:m7 c:maj7 d:7 e:m7 fis:hdim7 \bar "|."
}

fMajor = \chordmode {
  \key f \major
  f:maj7 g:m7 a:m7 bes:maj7 c:7 d:m7 e:hdim7 \bar "|."
}

\markup { \huge \bold \concat { C \super "△" } }
<<
  \new Staff = cMajor \cMajor
  \new ChordNames = cMajor \cMajor
>>

\markup { \huge \bold \concat { G \super "△" } }
<<
  \new Staff = gMajor \gMajor
  \new ChordNames = gMajor \gMajor
>>

\markup { \huge \bold \concat { F \super "△" } }
<<
  \new Staff = fMajor \fMajor
  \new ChordNames = fMajor \fMajor
>>

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\markup \huge \bold { Key types of 7th chords }

\relative {
  \clef treble
  \key c \major
  \time 4/4
  \textLengthOn
  <c'=' e g b>1-\markup { M7 }-\markup { MmM }-\markup { M3P5M7 } |
  <c e g bes>-\markup { D7 }-\markup { Mmm }-\markup { M3P5m7 } |
  <c ees g bes>-\markup { m7 }-\markup { mMm }-\markup { m3P5m7 } |
  <c ees ges bes>-\markup { sd7 }-\markup { mmM }-\markup { m3d5M7 } |
  <c=' ees ges beses>-\markup { d7 }-\markup { mmm }-\markup { m3d5d7 } \bar "|."
}

\relative {
  \clef treble
  \key g \major
  \time 4/4
  <g'=' b d fis>1-\markup M7 |
  <g b d f>-\markup D7 |
  <g bes d f>-\markup m7 |
  <g bes des f>-\markup sd7 |
  <g=' bes des fes>-\markup d7 \bar "|."
}

\relative {
  \clef treble
  \key f \major
  \time 4/4
  <f'=' a c e>1-\markup M7 |
  <f a c ees>-\markup D7 |
  <f aes c ees>-\markup m7 |
  <f aes ces ees>-\markup sd7 |
  <f=' aes ces eeses>-\markup d7 \bar "|."
}

\markup \huge \bold { \vspace #2 cM 7th chords }

\relative {
  \clef treble
  \key c \major
  \time 4/4
  <c'=' e g b>1-\markup M7 |
  <d f a c>-\markup m7 |
  <e g b d>-\markup m7 |
  <f a c e>-\markup M7 |
  <g b d f>-\markup D7 |
  <a c e g>-\markup m7 |
  <b=' d f a>-\markup sd7 \bar "|."
}

\markup \huge \bold { \vspace #2 gM 7th chords }

\relative {
  \clef treble
  \key g \major
  \time 4/4
  <g'=' b d fis>1-\markup M7 |
  <a c e g>-\markup m7 |
  <b d fis a>-\markup m7 |
  <c e g b>-\markup M7 |
  <d fis a c>-\markup D7 |
  <e g b d>-\markup m7 |
  <fis='' a c e>-\markup sd7 \bar "|."
}

\markup \huge \bold { \vspace #2 fM 7th chords }

\relative {
  \clef treble
  \key f \major
  \time 4/4
  <f'=' a c e>1-\markup M7 |
  <g bes d f>-\markup m7 |
  <a c e g>-\markup m7 |
  <bes d f a>-\markup M7 |
  <c e g bes>-\markup D7 |
  <d f a c>-\markup m7 |
  <e='' g bes d>-\markup sd7 \bar "|."
}

\markup \huge \bold { \vspace #2 am harmonic 7th chords }

\relative {
  \clef treble
  \key a \minor
  \time 4/4
  <a= c e gis>1-\markup mMM |
  <b d f a>-\markup sd7 |
  <c e gis b>-\markup MMm |
  <d f a c>-\markup m7 |
  <e gis b d>-\markup D7 |
  <f a c e>-\markup M7 |
  <gis=' b d f>-\markup d7 \bar "|."
}

\markup \huge \bold { \vspace #2 em harmonic 7th chords }

\relative {
  \clef treble
  \key e \minor
  \time 4/4
  <e'=' g b dis>1-\markup mMM |
  <fis a c e>-\markup sd7 |
  <g b dis fis>-\markup MMm |
  <a c e g>-\markup m7 |
  <b dis fis a>-\markup D7 |
  <c e g b>-\markup M7 |
  <dis='' fis a c>-\markup d7 \bar "|."
}

\markup \huge \bold { \vspace #2 dm harmonic 7th chords }

\relative {
  \clef treble
  \key d \minor
  \time 4/4
  <d'=' f a cis>1-\markup mMM |
  <e g bes d>-\markup sd7 |
  <f a cis e>-\markup MMm |
  <g bes d f>-\markup m7 |
  <a cis e g>-\markup D7 |
  <bes d f a>-\markup M7 |
  <cis='' e g bes>-\markup d7 \bar "|."
}

\markup \huge \bold { \vspace #2 am melodic 7th chords }

\relative {
  \clef treble
  \key a \minor
  \time 4/4
  <a= c e gis>1-\markup mMM |
  <b d fis a>-\markup m7 |
  <c e gis b>-\markup MMm |
  <d fis a c>-\markup D7 |
  <e gis b d>-\markup D7 |
  <fis a c e>-\markup sd7 |
  <gis=' b d fis>-\markup sd7 \bar "|."
}

\markup \huge \bold { \vspace #2 em melodic 7th chords }

\relative {
  \clef treble
  \key e \minor
  \time 4/4
  <e'=' g b dis>1-\markup mMM |
  <fis a cis e>-\markup m7 |
  <g b dis fis>-\markup MMm |
  <a cis e g>-\markup D7 |
  <b dis fis a>-\markup D7 |
  <cis e g b>-\markup sd7 |
  <dis='' fis a cis>-\markup sd7 \bar "|."
}

\markup \huge \bold { \vspace #2 dm melodic 7th chords }

\relative {
  \clef treble
  \key d \minor
  \time 4/4
  <d'=' f a cis>1-\markup mMM |
  <e g b d>-\markup m7 |
  <f a cis e>-\markup MMm |
  <g b d f>-\markup D7 |
  <a cis e g>-\markup D7 |
  <b d f a>-\markup sd7 |
  <cis='' e g b>-\markup sd7 \bar "|."
}
