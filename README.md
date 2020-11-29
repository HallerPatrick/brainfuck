# Brainfuck Interpreter in Go

My from scratch implementation of a
[Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) Interpreter written in Golang.

This 300 line project was meant to teach me some Go and maybe a little bit of interpreter crafting.

## Usage

````bash
go run main.go example/helloworld.b
````


This "simple" brainfuck program should print a good old `"Hello World!\n"`
```brainfuck
 +++++ +++++         		initialize counter (cell #0) to 10
[                   		set the next four cells to 70, 100, 30 and 10
	> +++++ ++          	add  7 to cell #1
	> +++++ +++++       	add 10 to cell #2
	> +++               	add  3 to cell #3
	> +                 	add  1 to cell #4
	<<<< -              	decrement counter (cell #0)
]               	

> ++ .              	print 'H'  (H = ASC (72))
> + .               	print 'e'  (e = ASC (101))
+++++ ++ .          	print 'l'
.                   	print 'l'
+++ .               	print 'o'

> ++ .              	print ' '

<< +++++ +++++ +++++ .  print 'W'
> .                 	print 'o'
+++ .               	print 'r'
----- - .           	print 'l'
----- --- .         	print 'd'
> + .               	print '!'
> .                 	print '\n'
```