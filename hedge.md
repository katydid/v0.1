---
layout: page
title: Bottom up Hedge Automata
---

Deterministic Finite Automata
-----------------------------

We all know DFAs (Deterministic Finite Automata), the base of regular expressions, they consist of a bunch of states and input symbols.
We begin in a start state, read an input symbol from a string and progress to a next state, etc. until the input string is exhausted.
If our final state is an accept state we have accepted the string otherwise not.

For example we could have a DFA that recognises any string which contains a substring <i>aab</i>.

<table>
	<tr><td><b>current</b></td><td><b>input</b></td><td><b>=</b></td><td><b>next</b></td></tr>
	<tr><td>start</td><td>a</td><td>=</td><td>first</td></tr>
	<tr><td>start</td><td>b</td><td>=</td><td>start</td></tr>
	<tr><td>first</td><td>a</td><td>=</td><td>second</td></tr>
	<tr><td>first</td><td>b</td><td>=</td><td>start</td></tr>
	<tr><td>second</td><td>a</td><td>=</td><td>second</td></tr>
	<tr><td>second</td><td>b</td><td>=</td><td>accept</td></tr>
	<tr><td>accept</td><td>_</td><td>=</td><td>accept</td></tr>
</table>

This is equivalent to the regular expression <i>.\*aab.\*</i>

![Image]({{ site.baseurl }}public/dfa.dot.gif)

The underscore character represents all other input symbols, which are not defined for the current state.
This means that, since we only have <i>a</i> and <i>b</i> in our input alphabet, we could rewrite:

<table>
	<tr><td>first</td><td>b</td><td>=</td><td>start</td></tr>
</table>

as:

<table>
	<tr><td>first</td><td>_</td><td>=</td><td>start</td></tr>
</table>

In this example we have assumed that <b>start</b> and <b>accept</b> are keywords.
We will continue with this assumption for <b>accept</b>, but we will learn how to indicate multiple start states for a hedge automaton.

Trees
-----

![Image]({{ site.baseurl }}public/tree.dot.gif)

DFAs recognise strings while Tree Automata recognise trees.
Hedge Automata are just Tree Automata for unranked trees.
I have not found any free and accessible explanations of tree automata.
I picked up most of the background from Joost Engelfriet's 1975 Lecture notes on tree automata and tree grammars otherwise the Handbook of Formal Languages have and will also come in handy.
I'll give an explanation by example on Hedge Automata in the next section.

First some tree terminology:

* a leaf node is a node without children, 
* a parent node is a node with children and 
* a root node is the node located at the root of the tree.

Hedge Automata
--------------

Hedge automata are basically just a lot of DFAs.
We start by turning each leaf node into an input symbol.
Next we start up an instance of a DFA on each parent node.
Each leaf then feeds its input symbol into its parent's DFA.
When all the leaves have fed themselves into their parent's DFA, the parent turns into an input symbol, corresponding to its final state, for its parent's DFA.
Finally we check whether the root DFA's final state is an accept state.

Lets make a hedge automaton which recognises a tree that has at least two nodes named <i>b</i> which each has at least one child named <i>e</i>.
This means that the example tree in the previous section will be recognised by our hedge automaton.

First we define input symbols for each type of leaf node:

<table>
	<tr><td><b>leaf node</b></td><td><b>=</b></td><td><b>input symbol</b></td></tr>
	<tr><td>c</td><td>=</td><td>x</td></tr>
	<tr><td>d</td><td>=</td><td>x</td></tr>
	<tr><td>e</td><td>=</td><td>e</td></tr>
</table>

Next we define DFAs for the parent nodes <i>a</i> and <i>b</i>:

<table>
	<tr><td colspan="4" align="center"><b>b: input alphabet = {e,x}</b></td></tr>
	<tr><td><b>current</b></td><td><b>input</b></td><td><b>=</b></td><td><b>next</b></td></tr>
	<tr><td>start</td><td>e</td><td>=</td><td>one</td></tr>
	<tr><td>start</td><td>_</td><td>=</td><td>start</td></tr>
	<tr><td>one</td><td>_</td><td>=</td><td>one</td></tr>
</table>

![Image]({{ site.baseurl }}public/hedgeb.dot.gif)

<table>
	<tr><td colspan="4" align="center"><b>a: input alphabet = {start, one}</b></td></tr>
	<tr><td><b>current</b></td><td><b>input</b></td><td><b>=</b></td><td><b>next</b></td></tr>
	<tr><td>start</td><td>one</td><td>=</td><td>one</td></tr>
	<tr><td>start</td><td>_</td><td>=</td><td>start</td></tr>
	<tr><td>one</td><td>one</td><td>=</td><td>accept</td></tr>
	<tr><td>one</td><td>_</td><td>=</td><td>one</td></tr>
	<tr><td>accept</td><td>_</td><td>=</td><td>accept</td></tr>
</table>

![Image]({{ site.baseurl }}public/hedgea.dot.gif)

Lets take our tree and run it through the hedge automaton.

![Image]({{ site.baseurl }}public/tree.dot.gif)

First the leaves are turned into input symbols.

![Image]({{ site.baseurl }}public/tree1.dot.gif)

Next an instance of an appropriate DFA is started on each parent node.

![Image]({{ site.baseurl }}public/tree2.dot.gif)

Now input symbols are fed into the the first DFA.
The DFA for node <i>b</i> is in state <i>start</i>, it reads <i>x</i> from its most left child, and it stays in state <i>start</i>.

![Image]({{ site.baseurl }}public/tree3.dot.gif)

Next it reads input symbol <i>e</i> and progresses to state <i>one</i>.

![Image]({{ site.baseurl }}public/tree4.dot.gif)

Finally it reads input symbol <i>x</i> and becomes an input symbol corresponding to its final state, <i>one</i>.

![Image]({{ site.baseurl }}public/tree5.dot.gif)

Now our root node reads input symbol <i>one</i> and goes from state <i>start</i> to state <i>one</i>.

![Image]({{ site.baseurl }}public/tree6.dot.gif)

Almost there. The other <i>b</i> node reads input symbol <i>e</i> and progresses from state <i>start</i> to state <i>one</i>.

![Image]({{ site.baseurl }}public/tree7.dot.gif)

Finally our root state reads input symbol <i>one</i> and progresses from state <i>one</i> to the <i>accept</i> state.

![Image]({{ site.baseurl }}public/tree8.dot.gif)