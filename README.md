# dontsleep

"Don't sleep" prevents HDD from going to sleep.

## Problem

For watching movies, I have a laptop connected to my TV
and the movies are stored on an external HDD. However,
this HDD goes to sleep very quickly when not used. When
I want to use it again, it requires several seconds to
spin up. How to keep this HDD awake, i.e. how to prevent it
from falling asleep?

## Solution

The idea is very simple. Read from or write to this HDD
every X seconds, thus it's used constantly and it cannot
go to sleep mode.

This simple program does exactly this. It writes an empty
file to the HDD every 5 minutes.

I measured my HDD and it shuts itself down in around
6 minutes and 20 seconds.

## Usage

Modify the value of the constant `fname`. Specify the path
of the file that should be written to.

Then build the EXE with `go build` and launch the
program. It will print some log on the screen.

## Improvement ideas

* The program could keep alive not only one but
several HDDs.

* The program could have a graphical user interface.

## Links

* [Prevent Hard Disk from going to Sleep in Windows 11/10](https://www.thewindowsclub.com/prevent-hard-drive-going-sleep-windows)
