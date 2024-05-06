The **solve** is king. Every action should serve the **solve** whether indirectly or indirectly.

What is a solve and how is it derived?

A solve is not:

- technologies and frameworks
- process
- meetings
  The implementation of the solve includes those things but they are not the solve.

The solve is the answer to the **question** that has the best tradeoffs.

Evaluation of tradeoffs comes with time and social connections. You need to have seen enough tactics fail and succeed in different environments to make projections.

Everything hinges upon these four questions:

- What needs solved?
- Who will solve it?
- Why should it be solved?
- When does it need to be solved by?

Sometimes other questions creep in, they only serve to distract you from the most useful path, it's best to not entertain them. Virtually every team can provide some rough answers for "who" and "when" - if they can't that team is dead in the water. I want to take that for granted and focus on "what" and "why".

## Good engineers ask "what needs solved?"

You begin your engineering career asking "what needs solved?" The problems begin scoped very small - there is little wiggle-room for possible implementations when you are tasked with fixing a typo on a webpage.

What needs solved?

The spelling of a word needs corrected.

The what determines the how. You don't even need to worry much about a "how" if your "what" is dialed in enough. An example - your company's eCommerce site goes down. One possible "what" you could try to solve is "The site is broken". It's not a very good "what", where do you even begin? If you were to drill down into the facts, check metrics, check logging, check updates that recently went out, etc. you could get a better what. You don't have much time, the site is totally offline and your company is loosing revenue by the minute. Even a quick what might be "The latest commit on master branch broke production." The immediate "how" is dumb simple: revert the commit and redeploy.

In a healthy organization there are two areas that determine an engineer's seniority. The first are soft skills. The second is **the probability that an engineer will deliver a great solve to a problem that has many possible solutions.** As a junior - you can't fail, your work is tightly scoped and defined, you simply need to execute, which comes down to technical skill. As you gain seniority, the bets get bigger. A senior engineer might work on a feature area taking weeks with a large amount of autonomy. Bad solves at that magnitude could harm the business. This is one of the reasons engineering interviews are so rigorous - a bad engineer could cause an extreme and recurring amount of damage to the business, even after their departure. A mediocre engineer could get by without harming the business but still leave huge profits on the table compared to a really exceptional engineer.

Technical skill is important but it's more like the absolute minimum bar to entry. The real value is in coming up with the best solve for the situation, which requires an excellent answer to "what needs solved?

## Great engineers ask "why should it be solved?

**The big O notation for not solving a problem is O(0).** It is theoretically the most optimal solve to any problem (until time travel is invented). Not doing something isn't always the best solution but it's the best solution far more often than people realize. Even when the big "what" that needs solved there will be thousands of sub-"whats" that usually have associated "whys".

Cross-examine yourself when you're doing something novel. You are an engineer after all so you must enjoy, or at the very least, are used to solving things. That means you are extremely biased and it's likely that your ego will cloud your judgement if you don't proactively work against it. "I enjoy writing object traversal algorithms" isn't a good answer to "Why should I write an algorithm to deep clone objects in JavaScript at my job?" When you put your ego aside in that situation - the answer to that "why" is "I shouldn't write that algorithm, I should use Lodash's `cloneDeep`." In case you're wondering - `cloneDeep` traverses an object pretty much like you'd expect with the interesting bit being a hashmap (confusingly named `stack` but not used like a stack...) that caches values to check for circular references.

Sometimes really interesting engineering problems do need solved on the job. Hopefully that happens a lot. Same principle as Jack LaLanne established for dieting, "if it tastes good, spit it out!", at the very least examine it before assuming it needs eaten.
