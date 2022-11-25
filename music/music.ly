\version "2.23.81"

chNameExceptions = {
  % % Triads
  % <c e g>-\markup \super "△"
  % <c ees g>-\markup \super "−"
  % <c e gis>-\markup \super "+"
  % % Seventh chords
  % <c e g b>-\markup \super "△7"
  % <c ees g bes>-\markup \super "−7"
  % <c e gis bes>-\markup { + \super 7 }
  % <c ees ges bes>-\markup \super "ø7"
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

chordmodifiers.hdm =
  #(lambda (pitches)
    (replace-step #{ ees' #} (replace-step #{ ges' #} pitches)))

cMaj = \chordmode {
  \key c \major
  c1:maj7 d:m7 e:m7 f:maj7 g:7 a:m7 b:hdm7 \bar "|."
}

cMinAlt = \chordmode {
  \key c \minor
  c1:m7 d:hdm7 ees:maj7 f2:m7 f:7 g:7 g:m7 aes1:maj7 bes2:7 b:dim7 \bar "|."
}

gMaj = \chordmode {
  \key g \major
  g1:maj7 a:m7 b:m7 c:maj7 d:7 e:m7 fis:hdm7 \bar "|."
}

gMinAlt = \chordmode {
  \key g \minor
  g1:m7 a:hdm7 bes:maj7 c2:m7 c:7 d:7 d:m7 ees1:maj7 f2:7 fis:dim7 \bar "|."
}

fMaj = \chordmode {
  \key f \major
  f1:maj7 g:m7 a:m7 bes:maj7 c:7 d:m7 e:hdm7 \bar "|."
}

fMinAlt = \chordmode {
  \key f \minor
  f1:m7 g:hdm7 aes:maj7 bes2:m7 bes:7 c:7 c:m7 des1:maj7 ees2:7 e2:dim7 \bar "|."
}

aMinHarm = \chordmode {
  \key a \minor
  a1:m7.7+ b:hdm7 c:maj7.5+ d:m7 e:7 f:maj7 gis:dim7 \bar "|."
}

eMinHarm = \chordmode {
  \key e \minor
  e1:m7.7+ fis:hdm7 g:maj7.5+ a:m7 b:7 c:maj7 dis:dim7 \bar "|."
}

dMinHarm = \chordmode {
  \key d \minor
  d1:m7.7+ e:hdm7 f:maj7.5+ g:m7 a:7 bes:maj7 cis:dim7 \bar "|."
}

aMinMel = \chordmode {
  \key a \minor
  a1:m7.7+ b:m7 c:maj7.5+ d:7 e:7 fis:hdm7 gis:hdm7 \bar "|."
}

eMinMel = \chordmode {
  \key e \minor
  e1:m7.7+ fis:m7 g:maj7.5+ a:7 b:7 cis:hdm7 dis:hdm7 \bar "|."
}

dMinMel = \chordmode {
  \key d \minor
  d1:m7.7+ e:m7 f:maj7.5+ g:7 a:7 b:hdm7 cis:hdm7 \bar "|."
}

\markup { \huge \bold { \concat { C \super "△" } / Cm } }
\new PianoStaff = cMajMin {
  <<
    \new Staff = cMaj \cMaj
    \new ChordNames = cMaj \cMaj
    \new Staff = cMinAlt \cMinAlt
    \new ChordNames = cMinAlt \cMinAlt
  >>
}

\markup { \huge \bold { \concat { G \super "△" } / Gm } }
\new PianoStaff = gMajMin {
  <<
    \new Staff = gMaj \gMaj
    \new ChordNames = gMaj \gMaj
    \new Staff = gMinAlt \gMinAlt
    \new ChordNames = gMinAlt \gMinAlt
  >>
}

\markup { \huge \bold { \concat { F \super "△" } / Fm } }
\new PianoStaff = fMajMin {
  <<
    \new Staff = fMaj \fMaj
    \new ChordNames = fMaj \fMaj
    \new Staff = fMinAlt \fMinAlt
    \new ChordNames = fMinAlt \fMinAlt
  >>
}

\markup { \huge \bold "Am harm." }
<<
  \new Staff = aMinHarm \aMinHarm
  \new ChordNames = aMinHarm \aMinHarm
>>

\markup { \huge \bold "Em harm." }
<<
  \new Staff = eMinHarm \eMinHarm
  \new ChordNames = eMinHarm \eMinHarm
>>

\markup { \huge \bold "Dm harm." }
<<
  \new Staff = dMinHarm \dMinHarm
  \new ChordNames = dMinHarm \dMinHarm
>>

\markup { \huge \bold "Am mel." }
<<
  \new Staff = aMinMel \aMinMel
  \new ChordNames = aMinMel \aMinMel
>>

\markup { \huge \bold "Em mel." }
<<
  \new Staff = eMinMel \eMinMel
  \new ChordNames = eMinMel \eMinMel
>>

\markup { \huge \bold "Dm mel." }
<<
  \new Staff = dMinMel \dMinMel
  \new ChordNames = dMinMel \dMinMel
>>

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

% \markup \huge \bold { Key types of 7th chords }

% \relative {
%   \clef treble
%   \key c \major
%   \time 4/4
%   \textLengthOn
%   <c'=' e g b>1-\markup { M7 }-\markup { MmM }-\markup { M3P5M7 } |
%   <c e g bes>-\markup { D7 }-\markup { Mmm }-\markup { M3P5m7 } |
%   <c ees g bes>-\markup { m7 }-\markup { mMm }-\markup { m3P5m7 } |
%   <c ees ges bes>-\markup { sd7 }-\markup { mmM }-\markup { m3d5M7 } |
%   <c=' ees ges beses>-\markup { d7 }-\markup { mmm }-\markup { m3d5d7 } \bar "|."
% }

% \relative {
%   \clef treble
%   \key g \major
%   \time 4/4
%   <g'=' b d fis>1-\markup M7 |
%   <g b d f>-\markup D7 |
%   <g bes d f>-\markup m7 |
%   <g bes des f>-\markup sd7 |
%   <g=' bes des fes>-\markup d7 \bar "|."
% }

% \relative {
%   \clef treble
%   \key f \major
%   \time 4/4
%   <f'=' a c e>1-\markup M7 |
%   <f a c ees>-\markup D7 |
%   <f aes c ees>-\markup m7 |
%   <f aes ces ees>-\markup sd7 |
%   <f=' aes ces eeses>-\markup d7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 cM 7th chords }

% \relative {
%   \clef treble
%   \key c \major
%   \time 4/4
%   <c'=' e g b>1-\markup M7 |
%   <d f a c>-\markup m7 |
%   <e g b d>-\markup m7 |
%   <f a c e>-\markup M7 |
%   <g b d f>-\markup D7 |
%   <a c e g>-\markup m7 |
%   <b=' d f a>-\markup sd7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 gM 7th chords }

% \relative {
%   \clef treble
%   \key g \major
%   \time 4/4
%   <g'=' b d fis>1-\markup M7 |
%   <a c e g>-\markup m7 |
%   <b d fis a>-\markup m7 |
%   <c e g b>-\markup M7 |
%   <d fis a c>-\markup D7 |
%   <e g b d>-\markup m7 |
%   <fis='' a c e>-\markup sd7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 fM 7th chords }

% \relative {
%   \clef treble
%   \key f \major
%   \time 4/4
%   <f'=' a c e>1-\markup M7 |
%   <g bes d f>-\markup m7 |
%   <a c e g>-\markup m7 |
%   <bes d f a>-\markup M7 |
%   <c e g bes>-\markup D7 |
%   <d f a c>-\markup m7 |
%   <e='' g bes d>-\markup sd7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 am harmonic 7th chords }

% \relative {
%   \clef treble
%   \key a \minor
%   \time 4/4
%   <a= c e gis>1-\markup mMM |
%   <b d f a>-\markup sd7 |
%   <c e gis b>-\markup MMm |
%   <d f a c>-\markup m7 |
%   <e gis b d>-\markup D7 |
%   <f a c e>-\markup M7 |
%   <gis=' b d f>-\markup d7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 em harmonic 7th chords }

% \relative {
%   \clef treble
%   \key e \minor
%   \time 4/4
%   <e'=' g b dis>1-\markup mMM |
%   <fis a c e>-\markup sd7 |
%   <g b dis fis>-\markup MMm |
%   <a c e g>-\markup m7 |
%   <b dis fis a>-\markup D7 |
%   <c e g b>-\markup M7 |
%   <dis='' fis a c>-\markup d7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 dm harmonic 7th chords }

% \relative {
%   \clef treble
%   \key d \minor
%   \time 4/4
%   <d'=' f a cis>1-\markup mMM |
%   <e g bes d>-\markup sd7 |
%   <f a cis e>-\markup MMm |
%   <g bes d f>-\markup m7 |
%   <a cis e g>-\markup D7 |
%   <bes d f a>-\markup M7 |
%   <cis='' e g bes>-\markup d7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 am melodic 7th chords }

% \relative {
%   \clef treble
%   \key a \minor
%   \time 4/4
%   <a= c e gis>1-\markup mMM |
%   <b d fis a>-\markup m7 |
%   <c e gis b>-\markup MMm |
%   <d fis a c>-\markup D7 |
%   <e gis b d>-\markup D7 |
%   <fis a c e>-\markup sd7 |
%   <gis=' b d fis>-\markup sd7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 em melodic 7th chords }

% \relative {
%   \clef treble
%   \key e \minor
%   \time 4/4
%   <e'=' g b dis>1-\markup mMM |
%   <fis a cis e>-\markup m7 |
%   <g b dis fis>-\markup MMm |
%   <a cis e g>-\markup D7 |
%   <b dis fis a>-\markup D7 |
%   <cis e g b>-\markup sd7 |
%   <dis='' fis a cis>-\markup sd7 \bar "|."
% }

% \markup \huge \bold { \vspace #2 dm melodic 7th chords }

% \relative {
%   \clef treble
%   \key d \minor
%   \time 4/4
%   <d'=' f a cis>1-\markup mMM |
%   <e g b d>-\markup m7 |
%   <f a cis e>-\markup MMm |
%   <g b d f>-\markup D7 |
%   <a cis e g>-\markup D7 |
%   <b d f a>-\markup sd7 |
%   <cis='' e g b>-\markup sd7 \bar "|."
% }
