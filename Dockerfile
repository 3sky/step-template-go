FROM debian

MAINTAINER 3sky <3sky@protonmail.com>

# Update package index and install go + git
RUN apt-get update && apt-get install -y \ 
    git \ 
    wget

#RUN mkdir -p /home/bitrise

RUN git clone https://github.com/3sky/step-template-go.git /home/bitrise

RUN wget https://github.com/bitrise-io/bitrise/releases/download/1.13.0/bitrise-Linux-x86_64 -O /home/bitrise/bitrise-Linux-x86_64

RUN chmod +x /home/bitrise/bitrise-Linux-x86_64

RUN /home/bitrise/bitrise-Linux-x86_64 setup


