# Pocketbase / Ember (w Tailwind!) Starter Template
I think that Pocketbase + Ember.js make a pretty good duo for full stack web dev. Here's a starter template for using them together.

## How it works
Well, there's probably opportunity to make this more robust, but in short, your ember app builds production builds into the `dist-prod` folder, which is then embedded (via go file embeds) into the pocketbase app and is served via a pocketbase route handler. I took some inspiration from how pocketbase serves their admin UI. 


## Why Ember? Isn't that a dead framework?
Ember isn't exactly the most popular front-end framework, but it is an absolute workhorse for helping quickly deliver fully featured SPAs. Ember is a great choice because its opinionated nature scales really well across teams... if you've seen one conventional ember app, you've seen them all. I always thought of it as the Rails of front-end JS.

Ember has strong opinions and was one of the first JS frameworks to heavily push the importance of routing and "routing first". I think from the beginning, Ember got a lot of things right with FE JS when JS frameworks were coming into style. In fact, Ember always pushed the need for frameworks and opinions, especially in "the wild west land of JS". We are now seeing the pendulum swing back with other JS libraries that are taking framework approaches. With all the freedom you get to design your app architecture, sometimes its nice to let a framework take over so you can focus on features.

From a personal perspective, I have had a lot of success in the last 10 years of my career using Ember. During the first half of my career, my incredible colleagues and I used it to quickly develop and ship many features on the [Heroku Dashboard](https://dashboard.heroku.com/). If you ever used the Heroku web UI, you for sure used features that my team and I developed. Most recently, we have had a lot of success with Ember at [Meroxa](https://meroxa.com/)... specifically with using it in conjunction with pocketbase.

I should probably write more here about this starter template, but it's late, so maybe another time 😴