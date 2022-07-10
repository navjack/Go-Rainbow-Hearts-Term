# Go-Rainbow-Hearts-Term
Silly thing that fills a 160x25 terminal with hearts. Result of noodling around with GitHub CoPilot.

## How to Use
Install Golang and just go build it. Then run it in your terminal of choice. It works on any system Go does. I'd like to hear what systems people have ran this on and how it performaned.

## Notes
Performance and results vary between OS and terminal application. Best results I've personally tested were in Hyper on MacOS. Worst was anything in Windows. Decent is the native Haiku terminal.

I'm not a programmer. Could this be faster and done better? Sure it could! I'm 90% sure though that the speed problem isn't in the fact that we have dumb loops drawing things and its more that I'm using control codes to do everything.

## MacOS (Hyper) (no speedup)
![ezgif-4-579bd7e24e](https://user-images.githubusercontent.com/7362750/178163569-93490c52-62ca-42d8-bf14-3c6eb8841e02.gif)

## Haiku (sped up 28x)
![ezgif-5-daee68c278](https://user-images.githubusercontent.com/7362750/178165646-1292f19b-4128-4185-ab82-ca7d62045e72.gif)

## Other Notes
I originally wanted to have a simple program that picked a random column on row 1 to draw a colored heart and then have it fall at a controlled tickrate to row 25 or on to another existing heart. I tried describing this to CoPilot as best I could but I wasn't able to get it just right. At one point I had an idea to somehow progmatically create all the variables using the Hearts struct but I couldn't figure that out. It makes sense to me to define 2000 Hearts (heart1, heart2...) using the Hearts struct for a simple 80x25 terminal. Each one would keep track of their location using the column and row variables. When a heart is falling, with each incriment of the row variable it would also check if the next location it would occupy was already occupied and if it was then it would stop moving. Loop this forever until the terminal was filled with the different colored Hearts.

Once I realized I couldn't figure that out I just kind of went with what I did here. At first it scrolled to the right but I got curious so I tried it the other way and I loved that it broke in such a way that it created that neat broken effect. So, I just looped it again and then back and now we have what I have here.
