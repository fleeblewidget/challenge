//Images and blocks of text used in the challenge
package main

const header string =`
       .     .                       *** **
                !      .           ._*.                       .
             - -*- -       .-'-.   !  !     .
    .    .      *       .-' .-. '-.!  !             .              .
               ***   .-' .-'   '-. '-.!    .
       *       ***.-' .-'         '-. '-.                   .
       *      ***$*.-'               '-. '-.     *
  *   ***     * ***     ___________     !-..!-.  *     *         *    *
  *   ***    **$** *   !__!__!__!__!    !    !  ***   ***    .   *   ***
 *** ****    * *****   !__!__!__!__!    !      .***-.-*** *     *** * #_
**********  * ****$ *  !__!__!__!__!    !-..--'*****   # '*-..---# ***
**** *****  * $** ***      .            !      *****     ***       ***
************ ***** ***-..-' -.._________!     *******    ***      *****
***********   .-#.-'           '-.-''-..!     *******   ****...     #
  # ''-.---''        Welcome To         '-....---#..--'****** ''-.---''-
              Electronic Letter Friend!!                  #

`
const mainMenuContent string = `
What would you like to do?

E)nter a code
W)rite to Santa
Q)uit
`

const mainMenuPic string =`

                                                _...._
                       \  _  /                .::o:::::.
                        (\o/)                .:::'''':o:.
                    ---  / \  ---            :o:_    _:::
                         >*<                 ':}_>()<_{:'
                        >0<@><             @    ''//\\''    @
                       >>>@<<*          @ #     //  \\     # @
                      >@>*<0<<<       __#_#____/'____'\____#_#_
                     >*>>@<<<@><<     [_________________________]
                    >@>>0<<<*<<@><     |=_- .-/\ /\ /\ /\--. =_-|
                   >*>>0<<@><<<@><<<    |-_= | \ \\ \\ \\ \ |-_=-|
                  >@>>*<<@><>*<<0<*<   |_=-=| / // // // / |_=-_|
    \*/          >0>>*<<@><>0><<*<@><<  |=_- |'-''-''-''-'  |=_=-|
___\\U//___     >*>>@><0<<*>>@><*<0<< | =_-| o          o |_==_|
|\\ | | \\|    >@>>0<*<<0>>@<<0<<<*<@><|=_- | !     (    ! |=-_=|
| \\| | _(UU)_ >((*))_>0><*<0><@><<<0<*<-,-=| !    ).    ! |-_-=|
|\ \| || / //||.*.*.*.|>>@<<*<<@>><0<<<((=_| ! __(:')__ ! |=_==|
|\\_|_|&&_// ||*.*.*.*|_\\db//__     (\_/)-|/^\=^=^^=^=/^\| _=_|
""""|'.'.'.|~~|.*.*.*|     ____|_   =('Y')=//,------------.
    |'.'.'.|   ^^^^^^|____|>>>>>>|  ( ~~~ )/(((((((())))))))
    ~~~~~~~~         '""""'------'  'w---w'  '------------'
`

const enterCodeContent string = `

                  YAY!! LET'S ENTER A CODE!!

          .--._.--.--.__.--.--.__.--.--.__.--.--._.--.
        _(_      _Y_      _Y_      _Y_      _Y_      _)_
       [___]    [___]    [___]    [___]    [___]    [___]
       /:' \    /:' \    /:' \    /:' \    /:' \    /:' \
      |::   |  |::   |  |::   |  |::   |  |::   |  |::   |
      \::.  /  \::.  /  \::.  /  \::.  /  \::.  /  \::.  /
  jgs  \::./    \::./    \::./    \::./    \::./    \::./
        '='      '='      '='      '='      '='      '='

Please enter your code:
`

const validCodeContent string = `

           CODE VALID!!!
    _  _       _______________________
  _(\)(/)_____/\_'_\ \/_______________\
 |'.)XX(=====/ /,---\*\----------------\
 |||'()_\____\// \ / \ \'               \
  '| |        Y   \___\ \________________\
    '|_________\  /   / /                /
                \/___/_/________________/
`

const writeSantaContent string = `
          __                                                      _.
 _---_.*~<('===          ,~~,         ,~~,         ,~~,           ,_)
(,    \ (__)=3--__._----_()'4__._----_()'4__._----_()'4__._,____.()'b
  \--------/-\  ~~(        ) ~~(        ) ~~(        )  ~~:       :'
   \_______|       (,_,,,_)     (,_,,,_)     (,_,,,_)     ;,,,,,,:
___I___I___I./     / /  \ \     / /  \ \     / /  \ \     / /  \ \

Alright, let's Write to Santa!

As you know, Santa will only answer letters which are written in the correct format. Electronic Letter Friend is a tool to help those with the correct knowledge to rewrite kids' letters in the right format so that they will get their toys on Christmas Eve.

Here is a letter from Kevin that needs re-writing. Please type the corrected letter below, ending with <END>.

Let's make some Christmas magic for poor, grammatically challenged little Kevin!

Letter begins:

"Dear Santa,

I want a whole bunch of stuff! I've been super good tho. I wanna bike, and also some sea monkeys, and a toy elephant you can ride on. And at least 100 colouring pencils. And mum says if I put that I like books you will bring me more stuff so please can I also have the Paw Patrol annual?

Kevin Michael Scott, 6 and 3/4

PS My house is at 107 Copse Lane Marston Oxford UK The world"
`

const successContent string = `

  CORRECT!!!!!!!

                      ()
                     vvvv
                   .-"  "-.
                 .'___     '.
                / /_/_/      \     ____
               | /_/_/        |  .%%%%%%.
               |              | /%/_/%%%%\_
                \         .::::.%%%%%%%%(_{}-o
                 '.     .::::::::%%%%%%%%/
                   '-._/:/_/::::::\%%%%%'
                   o-3_)::::::::::|""""
                       \::::::::::/
                        '::::::::'
                          ''''''     hjw

  You are now a successful Elf. The password is "Noel". CONGRATULATIONS!!!.
`

const goodbyeContent string = `
 	

                                    .
                        .       *
                              |         *
                   .    *         .            *
                                          /
                *      .       ,    *          .    *
                               ';.
                      *   -      ;:,   -    *     -   .
              .  -               ,:;:,
                     .          ,:;%;;:,           *
                          /    ::;%#%;::   *    .
                   *           ::;%#%;:'
                                ':%#%'  .   .,,.
                         *    .    #     .,sSSSSs
                                   #  .,sSSSSSSSS
                                .,sSSSSSSSSSSSS'
                           .,sSSSSSSSSSSSSSSSSSs,
                       .sSSSSSSSSSSSSSSSSSSSSSSSS
                       sSSSSSSSSSSSSSSSSSSSSSSSS'
                       'SSSSSSSSSSSSSSSSSSSSSSS'
                       sSSSS;nww;SSS;mwwwn;SSSSs
                       'SSS;nnw;sSSSs;wwwnnn;SSSs
                      .sSS;nnnw;SSSSS;wwwnnn%;SSS
                     .SSSS;nnnww;SSS;mwwwnnn%;SS'
                     SSSSS;nnnwwwmmmmmwwwnnn%%;
                     'SSS'%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
                        ;%%nnnwwwmmmmmwwwnnn%%;
      ,                 ;%%nnnwwwmmmmmwwwnnn%%;
   .,;;;,.,;            ;%%nnnwwwmmmmmwwwnnn%%;        ;     ;
     ';;;;,;;;,.,;      ;%%nnnwwwmmmmmnnnnnn%%;  ,    .;;,.,;;;.
      ;;;;;;;;,;;;;;,., ;%%nnnnnnnnn;,ooo,;n%%;   ;;;;;,;;,;;,;;;,.
      ;'   ';;;;;;,;;;;.;%%nnn;,ooo,;OOOOO;n%%;  .;;,;;;;;;;;;;,;;'',
            ;'  ';;;;,;;,...   OOOOO;'OOO'..,,;,;;,;;;''';;;''';;'
                  ';;'    ''''''OOO'OOooo' ''' ;'   '     '     '
                   '         ;,.,;, 'OOO'
                         ;,.;;';;;';,;.
                     ;,.;;';;;;;;;;;;;'
                   .,;;;;;;;;;;;;;;;'
                     ';'  ';'   ';;

    GOODBYE!!! AND MERRY CHRISTMAS!!!!!!!
`
