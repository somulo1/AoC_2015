<h2>--- Day 5: Doesn't He Have Intern-Elves For This? ---</h2>

Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

    It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
    It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
    It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

For example:

   <b> ugknbfddgicrmopn </b>is nice because it has at least three vowels (<b>u...i...o...</b>), a double letter (<b>...dd...</b>), and none of the disallowed substrings.
   </b> aaa </b>is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
    <b>jchzalrnumimnmhp </b>is naughty because it has no double letter.
    <b>haegwjzuvuyypxyu </b>is naughty because it contains the string <b>xy</b>.
    <b>dvszwmarrgswjxmb </b>is naughty because it contains only one vowel.

How many strings are nice?.............................

The first half of this puzzle is complete! It provides one gold star: *

<h2>--- Part Two ---</h2>

Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

    It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
    It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.

For example:

   <b> qjhvhtzxzqqjkmpb</b> is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
   <b> xxyxx</b> is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
    <b>uurcxstgmygtbstg </b>is naughty because it has a pair (tg) but no repeat with a single letter between them.
    <b>ieodomkazucvgmuy </b>is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.

How many strings are nice under these new rules?..............................

how to run the code:
`` go run . data.txt``
