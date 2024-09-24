
# Wand

Potter-ify your Linux experience!

Some basic Linux commands changed into Harry potter Spells!

# How to use

```sh
git clone https://github.com/fOSS-Community/wand

go mod tidy

go build -o bin

cd bin 

chmod +x wand

mv wand /usr/bin/
```

This will install wand globally on your pc.

### wand

```sh
$ wand
```

To get info about any spell, type, for eg- lumos:

```sh
$ wand lumos --help
```

# Spells Explained

The spells used in here are:
* [lumos](#lumos)
* [obliviate](#obliviate)
* [geminio](#geminio)
* [legilimens](#legilimens)
* [mobiliarbus](#mobiliarbus)
* [avada-kedavra](#avada-kedavra)
* [accio](#accio)
* [nox](#nox)
* [reducio](#reducio)

## <a id="lumos"></a>Lumos

Lumos is the incantation to a charm that can be used to produce a flash of bright white light from the tip of the wand.
This light can be thrown far off of the wand, illuminating the surrounding area for several minutes
after its initial casting before darkening once more.

Lumos shows the path in wizard world and so does the 'ls' command in linux world. 'ls' does directory listing.

So typing 'lumos' into terminal will display the files and folders in the current working directory.

Usage : `$ lumos`

**Cheers!**

## <a id="geminio"></a>Geminio

The Geminio Curse or Doubling Charm(Geminio) is a spell used to duplicate an object. It can also be used to bewitch an object into multiplying repeatedly when touched, though how one would produce the latter effect is still unknown.

It duplicates objects i.e. creates copies of its own and so does the 'cp' command in linux terminal. It copies files/directories in linux terminal.

Usage : `$ geminio /path/to/source /path/to/destination`

**Cheers!**

## <a id="obliviate"></a>Obliviate

Originating from "Harry Potter and the Chamber of Secrets", Obliviate is a term meaning 'forget'. Obliviate is a memory charm, resulting in the erasure of the recipient's memory. One's memory vanishes as soon as this charm is casted.

It makes one forgetful of things or memories and so does the 'clear' command in linux terminal. It clears the output screen or 'obliviates' it :P

Usage : `$ obliviate`

**Cheers!**

## <a id="legilimens"></a>Legilimens

Legilemency or legilimens is the act of magically navigating through many layers of a person's mind and correctly interpreting one's findings. Muggles often call this as 'mind-reading'.

It reads what is beneath the flesh of the person in wizard world, and somewhat 'cat' does the same in linux world. It reads the contents hidden beneath the flesh of a file :P

Usage : `$ legilimens filename`

**Cheers!**

## <a id="mobiliarbus"></a>Mobiliarbus

Mobiliarbus is the incantation to a charm used to levitate and move plants and trees, as well as the materials made of wood. The Latin term mobilis, meaning "movable" , and arbor means "tree".

It moves objects from one place to another, and so does the 'mv' command in linux world. It moves files and directories from one place to another.

Usage : `$ mobiliarbus /path/to/source /path/to/destination`

**Cheers!**

## <a id="avada-kedavra"></a>Avada-Kedavra

The Killing Curse (Avada Kedavra) is a tool of the Dark Arts and one of the three Unforgivable Curses. It is one of the most powerful and sinister spells known to wizardkind. When cast successfully on a living person or creature the curse causes instantaneous and painless death, without any signs of violence on body.

It 'deletes' the person or creature from the living world, poof! And so does the 'rm' command in linux world. 'rm' causes painless death of files and folders :P

Usage : `$ avada-kedavra filename`

**Cheers!**


## <a id="accio"></a>Accio

The Summoning Charm (Accio) was a charm that caused a target at a distance from the caster to levitate or fly over to them. This spell needs thought behind it, the object must be clear in the caster's mind before trying to summon.

This spell is one of the oldest spells known to wizarding society.

Accio 'gets' the object and so does 'wget' in the linux world.

Usage : `$ accio LINK`

**Cheers!**


## <a id="nox"></a>Nox


# License
This Package is licensed under MIT License.

# Author
[Kanishk Pachauri](https://github.com/Mr-Sunglasses)

Type : `$ wand SPELL_NAME --help` to get information regarding any spell.

**Cheers!**
