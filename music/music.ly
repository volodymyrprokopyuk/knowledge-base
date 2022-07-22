\version "2.23.10"

\markup \huge \bold { Key types of 7th chords }

\relative {
  \clef treble
  \key c \major
  \time 4/4
  \textLengthOn
  <c' e g b>1-\markup { M7 }-\markup { MmM }-\markup { M3P5M7 } |
  <c e g bes>-\markup { D7 }-\markup { Mmm }-\markup { M3P5m7 } |
  <c ees g bes>-\markup { m7 }-\markup { mMm }-\markup { m3P5m7 } |
  <c ees ges bes>-\markup { sd7 }-\markup { mmM }-\markup { m3d5M7 } |
  <c=' ees ges beses>-\markup { d7 }-\markup { mmm }-\markup { m3d5d7 } \bar "|."
}

\markup \huge \bold { \vspace #2 cM 7th chords }

\relative {
  \clef treble
  \key c \major
  \time 4/4
  <c' e g b>1-\markup M7 |
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
  a'1 |
}

\markup \huge \bold { \vspace #2 em harmonic 7th chords }

\relative {
  \clef treble
  \key e \minor
  \time 4/4
  e'1 |
}

\markup \huge \bold { \vspace #2 dm harmonic 7th chords }

\relative {
  \clef treble
  \key d \minor
  \time 4/4
  d'1 |
}

\markup \huge \bold { \vspace #2 am melodic 7th chords }

\relative {
  \clef treble
  \key a \minor
  \time 4/4
  a'1 |
}

\markup \huge \bold { \vspace #2 em melodic 7th chords }

\relative {
  \clef treble
  \key e \minor
  \time 4/4
  e'1 |
}

\markup \huge \bold { \vspace #2 dm melodic 7th chords }

\relative {
  \clef treble
  \key d \minor
  \time 4/4
  d'1 |
}
