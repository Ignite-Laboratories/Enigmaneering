# `E2S3 - Passages`
### `Alex Petz, Ignite Laboratories, July 2025`

---

### Binary Structures
So far, all of the work we've accomplished has been entirely focused around the `Phrase` structure - which is
just a high-level way of dynamically managing bits.  It's time for us to start bringing in some _logical_ structures 
of these bits.

    type Passage struct {
        Breadcrumbs Phrase
        Remainder Phrase
    }

Each breadcrumb marks a point on the `tracert` of the passage from logic into becoming, if you will.  