=begin

Decode the Morse code

In this kata you have to write a simple Morse code decoder.

While the Morse code is now mostly superceded by voice and digital data communication channels, 
it still has its use in some applications around the world.

The Morse code encodes every character as a sequence of "dots" and "dashes". For example, the letter A is coded as ·−, 
letter Q is coded as −−·−, and digit 1 is coded as ·−−−−. The Morse code is case-insensitive, traditionally capital letters are used. 

When the message is written in Morse code, a single space is used to separate the character codes and 3 spaces are used to separate words. 
For example, the message HEY JUDE in Morse code is ···· · −·−−   ·−−− ··− −·· ·.

NOTE: Extra spaces before or after the code have no meaning and should be ignored.

In addition to letters, digits and some punctuation, there are some special service codes, 
the most notorious of those is the international distress signal SOS (that was first issued by Titanic), that is coded as ···−−−···. 
These special codes are treated as single special characters, and usually are transmitted as separate words.

=end

MORSE_CODE = { 
  ".-" => "A", 
  "-..." => "B", 
  "-.-." => "C", 
  "-.." => "D", 
  "." => "E", 
  "..-." => "F",
  "--." => "G", 
  "...." => "H", 
  ".." => "I", 
  ".---" =>"J", 
  "-.-" => "K", 
  ".-.." => "L",
  "--" => "M", 
  "-." => "N",
  "---" => "O", 
  ".--." => "P", 
  "--.-" => "Q", 
  ".-." => "R",
  "..." => "S", 
  "-" => "T", 
  "..-" => "U", 
  "...-" => "V", 
  ".--" => "W", 
  "-..-" => "X", 
  "-.--" => "Y", 
  "--.." => "Z",
  "...---..." => "SOS",
  "-.-.--" => "!",
  ".-.-.-" => ".",
  ".----" => "1",
  "..---" => "2",
  "...--" => "3",
  "....-" => "4",
  "....." => "5",
  "-...." => "6", 
  "--..." => "7",
  "---.." => "8",
  "----." => "9",  
  "-----" => "0",   
}


def decode_morse(morse_code)
    morse_code.strip.split(/\s{3}/).map { |w| w.split(' ').map { |c| MORSE_CODE[c] }.join }.join(' ')
end
