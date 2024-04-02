# 🌟 Stegologin

**Stegologin** is a proof of concept for implementing login functionality that uses steganography to hide credentials during HTTP transfer.

## 🕵️‍♂️ Why Hide the Credentials?

The idea came to me after watching a video about secret CIA websites camouflaged as Iranian soccer club homepages. These sites allowed secret agents to interact with Langley. What did all these websites have in common? A textbox at the bottom where agents would enter their secret code to gain access. So, I asked myself, why keep it in plain sight when you could make it a little more sci-fi and Bond-like?

## 🚀 Implementation

The HTTP server is written in Go using [Fiber](https://github.com/gofiber/fiber).

For steganography, I use [this](https://github.com/auyer/steganography) awesome library written in pure Go.

Certainly! Here's the updated section for your README:

## 🌟 Practical Use Case

Imagine creating a simple forum where users can register and upload profile pictures. But let's add a twist: secret agents are among the users, and they need to hide their credentials within those profile pictures. 🕵️‍♂️

1. **User Registration**: Regular users sign up with their usual usernames and passwords. Nothing out of the ordinary here.

2. **Secret Agent Registration**: When secret agents register, they follow the same process, but with an additional step. After entering their credentials, they upload a profile picture. This seemingly innocent image becomes their secret key.

3. **Steganography Magic**: Using steganography techniques, the server embeds the agent's credentials (perhaps an access code or passphrase) directly into the profile picture. The hidden information is imperceptible to the naked eye.

4. **Authentication**: When the agent logs in, the server extracts the hidden credentials from their profile picture. If the extracted data matches what's stored on the server, access is granted. Mission accomplished!

Remember, this is just a fun concept, but it highlights the creative possibilities of steganography in everyday applications. 🚀
