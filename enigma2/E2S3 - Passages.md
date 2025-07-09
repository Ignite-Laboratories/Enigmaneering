# `E2S3 - Passages`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### Binary Structures
So far, all of the work we've accomplished has been entirely focused around the `Phrase` structure - which is
just a high-level way of dynamically managing bits.  But what if you want to hold several varying width phrases
together?  We _could_ write a structure similar to a map with a named set of phrases, but the interface to that
in code would be _quite clunky_ as you'd have to access phrases by a string accessor like `passage["breadcrumbs"]`.
Instead, I believe this is the point where I get to fork off _my_ logical structures and teach _when_ an abstraction
is helpful!  In this case, we have a unique pairing of phrases that we will need to reference _**a lot.**_

    // Passage represents a breadcrumb path
    type Passage struct {
        Breadcrumbs Phrase
        Remainder Phrase
    }

Each breadcrumb marks a point on the `tracert` of the passage from logic into becoming, if you will.  