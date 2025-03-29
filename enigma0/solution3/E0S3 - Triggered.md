# `E0S3 - Triggered`
### `Alex Petz, Ignite Laboratories, March 2025`

---

### What if you want to fire something _once?_

That's easy!  Except, your actions now are infected with a `Context` parameter...  How can you fire
it off just once, and on _your terms!?_  Well - _exactly as before!_  

In this example, the data from `stdin` becomes the trigger for stimulating a readout of the next impulse moment -

    func init() {
        go core.Impulse.Spark() // Make it so
    }
    
    func main() {
        for core.Alive {
            // Press the enter key to read from stdin
            _, _ = fmt.Scanln()

            // Trigger a stimulation
            core.Impulse.Trigger(PrintParity, condition.Always, true)
        }
    }
    
    func PrintParity(ctx core.Context) {
        fmt.Printf("Impulse moment %v\n", ctx.Moment)
    }

Here we are simply reading from stdin and then triggering an impulse if it gives us any data. Obviously,
this is a highly contrived example, but it demonstrates the _reuse_ of temporal methods by non-temporal
threads.  Here, the main execution thread is able to stimulate the PrintParity function without any
fancy loops.

Even though it's a single activation, it still requires a potential to operate.  In most cases, `when.Always`
is a perfectly acceptable potential to utilize.

The last parameter is `async` and indicates if the triggered activation should operate asynchronously or
in a blocking fashion.

### What if you want to trigger the neuron directly?
That's also acceptable!  The neuron type provides a `Trigger(async)` function that allows it to directly
call `Trigger` on the engine it was created from.