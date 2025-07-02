# `E1S1 - Binary Subdivision`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### Getting Intimate With Binary
The next part to recognize is the difference between numerical and logical binary data.  For synthesis, we work
with the numerical representation of data while logically managing the leading zeros.  That will make a lot
more sense as we proceed forward - for now, just know that numeric equivalence does _not_ imply logical equivalence.
A single missing zero could entirely destroy the logical structure of data - fun! =)

      ⬐ Full bit width   ⬐ Equivalent numeric value
    [ 0 0 1 0 1 0 1 0 ] (42)  Logical form
        [ 1 0 1 0 1 0 ] (42)  Numeric form
              ⬑ Truncated bit width

Here, the logical form is a _byte_ - but that's not a hard requirement, just a universal standard.

This is what we ultimately will be exploiting to facilitate binary synthesis!  Some values can be stored in
_far less bits_ than they're logically stored at, relative to known points in the index.  For example, if 
walking down from the dark side the numeric bit length shrinks exponentially.  If one instead counted from 
a known point in the index, the numeric length changes.  To highlight this, one can _synthesize_ points using 
a _bit pattern_ repeated across the index.

Let's take an 11-bit index and subdivide it into eight regions using a note (3-bit) pattern:


    Note Subdivision of an 11 bit Index:
 
         Pattern        Synthesized          Value    Delta   
     (0) [ 0 0 0 ] [ 0 0 0 0 0 0 0 0 0 0 0 ] (   0  ) + 292
     (1) [ 0 0 1 ] [ 0 0 1 0 0 1 0 0 1 0 0 ] (  292 ) + 293
     (2) [ 0 1 0 ] [ 0 1 0 0 1 0 0 1 0 0 1 ] (  585 ) + 292
     (3) [ 0 1 1 ] [ 0 1 1 0 1 1 0 1 1 0 1 ] (  877 ) + 293
     (4) [ 1 0 0 ] [ 1 0 0 1 0 0 1 0 0 1 0 ] ( 1170 ) + 292
     (5) [ 1 0 1 ] [ 1 0 1 1 0 1 1 0 1 1 0 ] ( 1462 ) + 293
     (6) [ 1 1 0 ] [ 1 1 0 1 1 0 1 1 0 1 1 ] ( 1755 ) + 292
     (7) [ 1 1 1 ] [ 1 1 1 1 1 1 1 1 1 1 1 ] ( 2047 )

Literally any width index can be evenly subdivided by the bit width of your pattern simply by repeating
the pattern across that bit width.  If the value that point represents is a floating point number, binary
truncates it to the closest whole number _naturally._  

It truly is the most beautiful counting system in existence =)

### Why?
Subdivision, by itself, isn't anything special - in fact, we don't use it at all in the process!  However,
it provides a way to visualize counting to a number from a _synthetic point._  For the synthesis process
we will specifically be counting from the _midpoint,_ rather than _0._

The solution here is a primitive demo to prove that binary follows these subdivision rules for any index width. 