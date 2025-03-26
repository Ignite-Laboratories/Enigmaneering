# `E0S2 - Neural Muting`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### What's the point of muting?

Well, because a potential tells if a neuron _should_ activate when impulsed - but muting indicates if
the neuron should even _test_ it's potential.  This is a very powerful feature when combined with
a neural _reference_ -

    var printer = core.Impulse.Stimulate(PrintParity, condition.Always, true)
    
    func main() {
        core.Impulse.Loop(Toggle, condition.Always, false)
        core.Impulse.MaxFrequency = 4
        core.Impulse.Spark()
    }

Here the `printer` global is a `*Neuron` type, which provides rudimentary access to the underlying
action potential.  In our above configuration, the `Toggle` function is looped -

    func Toggle(ctx core.Context) {
        if printer.Muted {
            fmt.Printf("[%d] Un-muting\n", ctx.Beat)
        } else {
            fmt.Printf("[%d] Muting\n", ctx.Beat)
        }
        printer.Muted = !printer.Muted
        time.Sleep(time.Second * 2)
    }

In here, the neuron can directly be _suppressed_ and _unsuppressed_ from activation without affecting
its underlying potential.  This allows neurons to lie _dormant_ until another system un-suppresses it,
and demonstrates the beginnings of intra-neural _signaling_.

Hey, look at that - the examples are getting less complex!  That's a wonderful sign =)