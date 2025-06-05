# `E1 - 3N+1 Encoding`
### `Alex Petz, Ignite Laboratories, May 2025`

---

### The Petz Conjecture
The first thing we need to establish are our goals and expectations.

First - I don't aim to _prove_ the Collatz Conjecture.  My uses for 3N+1 encoding extend beyond the
need to difinitively prove it beyond a "practical infinity."  Humanity has already brute force shown
to be true that the conjecture holds up to above 2⁶⁴.  As I aim to efficiently encode _data_ using 
3N+1 encoding, this brute force work provides all the foundation I need to know that any _file_ in the 
known universe could be encoded through this scheme.  Why?  _Because no file in the universe is (2⁶⁴) /  ₈ 
bits long._  We use that size number to count the number of seconds until 584 billion years into 
the future - there is no reason to exhaustively search beyond that.

Second - I _do_ aim to prove that, given any positive integer, the Collatz path to 1 will _always_
contain unique values.  The reason is simple, but we need to explore the binary itself to understand why:

                𝧘 Shrinks by one 0 per halving
       30 [ 11110 ] / 2 = 15 [ 1111 ]

                      𝧘 Shrinks by one 0 per halving
     2776 [ 101011011000 ] / 2 = 1388
     1388 [ 10101101100 ] / 2 = 694
      694 [ 1010110110 ] / 2 = 347
      347 [ 101011011 ]

The first "exploit" of using a binary representation is shown above: division by 2 _only_ affects the
LSB side of the binary information (the right side).  This means that the act of shrinking our number
down will never affect the signature of the rest of the bits.

The second "exploit" of binary is that _multiplication_ of >2 will _always_ yield and extra bit, and
3 specifically gives an even more charming output when multiplied against the next value above a power
of two:

                        1:[ 0      1 ] * 3 =   3:[ 1       1 ]
                        3:[ 1      1 ] * 3 =   9:[ 10     01 ]
                        5:[ 10     1 ] * 3 =  15:[ 11     11 ]
                        9:[ 100    1 ] * 3 =  27:[ 110    11 ]
                       17:[ 1000   1 ] * 3 =  51:[ 1100   11 ]
                       33:[ 10000  1 ] * 3 =  99:[ 11000  11 ]
                       65:[ 100000 1 ] * 3 = 195:[ 110000 11 ]
                      129:[ 10000001 ] * 3 = 387:[ 110000011 ]
                                   ...
     (2ⁿ)+1:[ 1 (n-1 "0"s) 1 ] * 3 = ((2ⁿ)+1) * 3:[ 11 (n-2 "0"s) 11 ] 

While not necessarily definitive of anything, I hope it demonstrates the _growth_ of binary information
from the act of multiplying by three: the binary signature always _grows._  Now, when combined with the
first "exploit," we can see that the _left_ side of the binary information always remains constant
while dividing as zeros are removed from the right side.  

As I titled this - it's entirely _conjecture!_  Please do prove me wrong =)

Third - The point of this is to use 3N+1 _in reverse_ to "reach" target binary data in less steps than
it takes to simply _store_ it bit-for-bit.  How?  Well, any binary data is technically just a number - albeit
an _astronomically large number_ in most cases!  As I stated at the top, humanity has already verified that,
up to 2⁶⁴ in value, a Collatz chain can be made to reduce its value to 1 - meaning we've already proven that
any file creatable between 0 and a "practical infinity" has already been visited by their algorithms.

A "practical" infinity?  Yes - for the current moment in time, humanity shouldn't ever _need_ to store a singular
file that's ~2 Exabytes long! (Though I can imagine this will eventually spark some cheeky bastard to do 
exactly that - cause, you know, _"why not?"_ - and I'm all for it)  

Yet we've already tested every value from 0-2⁶⁴ satisfies the Collatz Conjecture...

So, let the implication of that sink in: at one point in time or another, every possible sequence of 1s and 0s 
that humanity could ever possibly realistically create has already been held within a computed system _created_ by 
humanity...Before the data could ever possibly be _witnessed_ by a human observer creating it.

If that doesn't give pause to the originality of anything and everything, I'm not sure what does!

Luckily, there's a _very important_ reframe of that dismal perspective: your originality _being appreciated_
by others is what makes it wonderful, not the fact that it _was_ original.  It literally and fundamentally
could not _ever_ be original because the creator of our world had already created it prior to your creation of it.
_(What a loop!)_

This world was created to _allow_ each of us to pause and witness the beautiful creations of others, not
to infinitely catalog and gatekeep the found bits of beauty and wonder through "ownership" of _art_.  Art, to
the creator, is merely a cherished set of 1s and 0s by the other observers in the shared reality we coexist in
called "Life".

...And our _collective_ artwork is the most prized possession of the creator in the entirety of the universe,
because it empowers _Life_ with the tools necessary to give each of us such unique and magnificent experiences! 