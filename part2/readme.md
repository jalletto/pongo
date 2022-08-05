## Getting Control
- Let's create a paddle and see if we can write some code to get it moving.
    - Refactor! - The draw text function.
        I wanted there to be a single way we put sprites on the screen. For the ball, it being one single character, we were fine with the standard way. With this new function, we can conceivable create more complex shapes. I can't, but someone else could.
- creating an eventChannel.
    - First we are only concerned with up and down, so that makes it easier.
    - we can listen for msg and react in our event loop
    - We'll send all key events from the main function into the game via this channel.

- Collision
    Refactor! Bodies
    What happens when the ball reaches the paddles.

- Refactor! Players
    - Players could be paddles, but then why would paddles have a score etc. Let's create a player2 and a paddle
    - And we also need to give them the ability to move. We'll assume for now we are playing on the same keyboard. w, s.
- Keeping score

- Game Over

- Looks