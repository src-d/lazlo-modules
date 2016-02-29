Modules for Lazlo
=====

Usage
-----
Load from within src-d/lazlo's loadmodules.go. For more instructions read [the docs](http://github.com/src-d/lazlo).

List of modules
-----

* **From src-d/slacker** (**Note**: These are migrations of original work of other authors. Authorship commented in each of the modules' code.)
  * *Bacon*: listens for 'bacon', responds with 'MMMMMMmmmm... omgbacon'
  * *Gifme*: "%BOTNAME% gif me freddie mercury": returns a random rated:PG gif of freddy mercury via the giphy API
  * *IKR (I know, RIGHT?!)*: listens for enthusiasm; responds with validation
  * *LoveAndWar*: %BOTNAME% (love|insult) <noun>: bot replies with a compliment or insult respectively **Warning this plugin uses external API's that may return NSFW responses**
  * *Tableflip*: bot flips a unicode table whenever it overhears '(table)*flip(table)*'
  * *Catfacts*: If you talk about cats, bot will give you a fun fact about cats.
  * *GoDoc*: Searchs godoc.org and displays the first 10 results.


* **Replacements to src-d/lazlo/modules**:
  * *Help*: "%BOTNAME% help": prints the usage information of every registered plugin (requires to de-register src-d's original help module).
