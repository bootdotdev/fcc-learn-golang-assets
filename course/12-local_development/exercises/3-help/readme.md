# About this chapter

Our courses can be completed *almost* completely in the browser, but there are some things that you need to do on your local machine for the sake of learning!

For the majority of this chapter, coding in the browser won't be useful because you need to learn how Go code is organized in a local development environment. We'll be explaining *how* to do this and quizzing you on it.

This section will almost entirely be quiz-style, with very little coding in the browser. As usual, if you get lost please use our discord channel as a resource:

[Discord Community](https://boot.dev/community)

## An example project to follow along with

If you get lost during this chapter, refer to this `src` folder of our "Social Media Server in Go" project that can be [found here](https://github.com/bootdotdev/projects/tree/main/projects/social-media-backend-golang/10-posts_endpoints/src). It will serve as a basic example of many of the concepts in this chapter.

## Unix

This guide will assume you are on a Unix environment like Linux or Mac. If you're on Windows you may have to do just a *bit* of Google-ing or ask in Discord to figure out how some commands translate to Windows.

If you are on Windows, I'd optionally recommend checking out [WSL (Windows Subsystem for Linux)](https://docs.microsoft.com/en-us/windows/wsl/install) so that you can work in a Unix environment on your local machine.

## Download Go locally

I typically recommend one of two ways:

* [Official Download](https://golang.org/doc/install)
* [Webi Downloader](https://webinstall.dev/golang/)

Make sure to use at least version `1.20`. This can be verified after installation by typing:

```bash
go version
```
