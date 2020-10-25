FROM golang:latest

ENV GO111MODULE=on
ENV GOLANG_VERSION=1.15.3

WORKDIR /go/src/backend

COPY /backend/go.mod /backend/go.sum /go/src/backend/

RUN apt-get update -y \ 
    && apt-get install -y vim \
#     && mkdir -p ~/.vim/autoload ~/.vim/bundle \
#     && curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim \
#     && echo 'execute pathogen#infect()' >> ~/.vimrc \
#     && echo 'syntax on' >> ~/.vimrc \
#     && echo 'filetype plugin indent on' >> ~/.vimrc \
#     && echo 'let g:go_null_module_warning = 0' >> ~/.vimrc \
#     && echo 'set clipboard=unnamed,autoselect' >> ~/.vimrc \
#     && git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go \