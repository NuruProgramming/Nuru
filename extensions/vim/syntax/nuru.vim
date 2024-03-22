" Sintaksia ya nuru kwenye programu ya "vim"
" Lugha: Nuru

" Maneno tengwa
syntax keyword nuruKeyword unda pakeji rudisha vunja endelea tupu
syntax keyword nuruType fanya
syntax keyword nuruBool kweli sikweli
syntax keyword nuruConditional kama sivyo au
syntax match nuruComparision /[!\|<>]/
syntax keyword nuruLoop ktk while badili
syntax keyword nuruLabel ikiwa kawaida

" Nambari
syntax match nuruInt '[+-]\d\+' contained display
syntax match nuruFloat '[+-]\d+\.\d*' contained display

" Viendeshaji
syntax match nuruAssignment '='
syntax match nuruLogicalOP /[\&!|]/

" Vitendakazi 
syntax keyword nuruFunction andika aina jaza fungua

" Tungo
syntax region nuruString start=/"/ skip=/\\"/ end=/"/
syntax region nuruString start=/'/ skip=/\\'/ end=/'/

" Maoni
syntax match nuruComment "//.*"
syntax region nuruComment start="/\*" end="\*/"

" Fafanua sintaksia
let b:current_syntax = "nuru"

highlight def link nuruComment Comment
highlight def link nuruBool Boolean
highlight def link nuruFunction Function
highlight def link nuruComparision Conditional
highlight def link nuruConditional Conditional
highlight def link nuruKeyword Keyword
highlight def link nuruString String
highlight def link nuruVariable Identifier
highlight def link nuruLoop Repeat
highlight def link nuruInt Number
highlight def link nuruFloat Float
highlight def link nuruAssignment Operator
highlight def link nuruLogicalOP Operator
highlight def link nuruAriOP Operator
highlight def link nuruType Type
highlight def link nuruLabel Label

