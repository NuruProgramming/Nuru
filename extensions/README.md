# Nuru Extensions For Various Editors

## [VSCODE](./vscode/)

Nuru syntax highlighting on VSCode

## [VIM](./vim)

The file contained herein has a basic syntax highlight for vim.
The file should be saved in `$HOME/.vim/syntax/nuru.vim`.
You should add the following line to your `.vimrc` or the appropriate location:

```vim
au BufRead,BufNewFile *.nr set filetype=nuru
```

Only basic syntax highlighting is provided by the script.
