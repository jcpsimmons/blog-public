There are several very odd takes on technical leadership roles cropping up lately on blog posts, Hacker News, and LinkedIn. The original version of this post was a direct response to one of these essays describing the transition from IC track to management in software engineering. Reading that essay was what inspired me to attempt to bring some clarity to some points about tech leadership. I’ve since revised the essay to be more general due to the pervasiveness of strange viewpoints around this topic that have been cropping up lately.

A note before hopping in.

I am not trying to boohoo leadership responsibilities. The point I am making is that leadership roles are not some panacea that instantly makes life easier. They simply trade one problem set for another one so you should pick the problems that you have a natural aptitude for. It seems lately that there has been a lot of, frankly, unearned prestige assigned to these roles. Although I am not a believer in so-called "servant leadership" I do see the importance of tech leaders not only shaping the direction of effort but also empowering their reports so that they can set and achieve good career goals.

## Leadership is Where the Money is

Ever since COVID there has been a myopic focus among tech workers on maximizing compensation. I believe that everyone should negotiate assertively on their behalf, acquire negotiating power through being skilled at their job, and make a good living commensurate with the effort they put in, however, making a very high salary doesn't mean much if you're miserable. I know this firsthand, I quit a very prestigious and high-paying job after just five months because I was so astoundingly bored.

### The Pay Difference

Ignoring the fact that you'll be miserable if you make career decisions solely based on money, you won't necessarily earn more money in leadership roles. Pay in tech is primarily determined by the following factors:

- Stock appreciation/depreciation during the employee's tenure at a company
- Technical skill - the highest-paying companies only hire the best talent.
- Geographic location
- Negotiation skills and leverage during the offer stage

It really is that simple. An engineering manager at an early-stage startup might only make 120k while a mid-level software engineer at Meta makes about 300k.

### Parallel Tracks

All else being equal - comparing level-to-level at a specific company doesn't yield a much clearer picture, e.g. engineering management track typically starts at the same ranking level as a senior IC, level 5. Sometimes EMs get paid more than their equally leveled counterparts but sometimes the inverse is true (see levels.fyi salary data for Meta and Google to see an example of each). At any rate - if the pay is not the same, it is in the same neighborhood.

In the domain of engineering, levels proceed beyond the “fork in the road” at level 5. ICs never have to take on direct reports and can keep progressing in parallel to their management track colleagues both in terms of leveling and pay band.

## I Don’t Enjoy Coding

If you don't enjoy coding you should reconsider tech as a career. Perhaps you're not dissatisfied with coding but with the stack you're working in or even the company you're employed at. At any rate, it is essential to deeply connect with a love of coding before considering specialization - whether it's in engineering management, or the higher levels of IC engineering. Specialization of any kind is not a "fallback plan". If you go into engineering management without a solid background in the fundamentals, you won’t grok what’s going on at a deep enough level to become excellent, and your reports won’t respect you.

## The Best Engineers Should Become Managers

This is an inversion of the proper way of doing things. Two things happen when an organization moves its best engineers into management solely because they are the best engineers. You get:

- Engineering managers who have an unknown chance of success in their new positions
- A poorly built product because you gutted the department of its best engineers

Can an exceptional engineer become a great engineering manager? Sure, in the same sense that a coin flip might come up heads. The real question is why would you index on coding ability as _the skill_ that predicts success as an engineering manager? Above a certain level of technical competence, those two skill sets are largely independent variables.

So what can go wrong when the best engineer is moved into the management track? All sorts of uncomfortable things if they don't already have, or are unable to develop the leadership, people, and product skills required to succeed. One of the common smells of this is treating direct reports like computer programs. A good example that I've seen ebb and flow in trendiness during my career is treating lines of code as a useful metric for evaluating employees. In computer-world it's easy to say that "big number is good number". In people-world the picture is far more complex and qualitative. This approach fails to see that the one-line code change which was the result of weeks of research and endless conversation that successfully reduced a bunch of customer friction in the sales funnel and increased sales significantly probably provides more value than a 20,000 line giga-PR of automatically generated test data. Looks like CI-TEST-GEN-BOT is getting a big promotion this year!

## It's Not a Fork in the Road

Going into management or advanced IC roles is often perceived as a "fork in the road" while instead there is a wide common area between these two roles. I believe the roots of this myth have their origins in those new to the field primarily forming their opinions based on what they read on the internet instead of first-hand experience in the industry - this seems especially common among engineers who began their careers as remote employees.

In reality - no remotely well-functioning organization will hire a staff engineer who cannot lead and communicate. There is a strange perception, usually held by juniors and people on the business side of things, that staff engineers are holed up in a dark room furiously typing code every day for 12 hours. Maybe this is true somewhere out there but I've never experienced it myself nor seen it. Staff engineers, on average, usually code _less_ than senior engineers. A good deal of their time is spent in meetings, writing proposal documents, planning work, and getting buy-in for architecture decisions.

Conversely - nobody will hire an engineering manager who cannot build systems. Of course, sometimes these individuals do slip through the cracks, but their long-term prospects are not good. I saw this happen years ago to a friend's team. Their manager resigned to take another job so a new hire was made to backfill them. The person they selected was a decent talker but their technical knowledge was woefully out of date. As they ramped up and attempted to lead, the team ate them alive. Whenever architecture decisions needed to be made, they were made based on the tools that the manager was familiar with instead of the tools with the best tradeoffs. When questioned about these decisions, they obviously couldn't mount a defense for "why old-framework-X instead of new-framework-Y that is more performant?" because they were unwilling to learn.

Of course, skills can atrophy over time, but engineering managers should always keep current on the latest trends in their discipline. There are two great ways I've found to do this. One is to build small personal projects in the evenings. My latest is a small program that grabs a quote from a list that I keep, turns it into speech via OpenAI's text-to-speech API, and then sends it to me on Discord. It does this right around the time I wake up for the day.

Another great way to keep some fresh knowledge of new tech is to build small dev tools for your team. The latest I built was an agile planning poker CLI. The trick is to find something useful that won't be blocking if you can't work on it for a week or two.

## Alternate Approach

Grind hard at the early stages of your career. Don't say no to any work that comes across your desk - especially if you're young, you have more energy to burn. Don't worry too much about focus at first - concern yourself with getting a wide breadth of experience. Make sure to spend some time working at a small startup where you have to do more than your titled role indicates. Entitlement thinking disseminated through social media will try to convince you that you are a sucker for going above and beyond. This is extremely short-sighted. Of course, doing more work benefits the organization you work for but it is also the best long-term investment you can make in your career since you are effectively learning while being paid. As someone who spent way too much money on grad school, I want to emphasize that getting paid to learn is a good deal!

So what does going above and beyond look like? Follow your nose and sense where you can provide some extra value. Maybe the product manager on your team needs help with requirements gathering. Perhaps your marketing team needs better analytics data from your core app. There are opportunities like this everywhere, but more so at smaller startups, hence my insistence that you should at least spend a little time working in those environments.

Conversely, don't stay in an organization past the point at which you can grow, atrophy is one of the few things that can seriously derail your career in its early stages. This blunder can be extremely hard to overcome later on since you will have frittered away your most productive learning years not learning. **In general, you should not care too much about what they are paying you, optimize instead for building the best pedigree possible. Identify good mentors during this time - don't simply develop a rapport with them, become useful to them. Usefulness and competence are increasingly rare and therefore highly valued. Flattery is cheap and plentiful.**

You need two things to excel, to be good at what you do, and to enjoy what you do.

### Being Good at Something

Once you've racked up some years of wins and losses, pay attention to feedback from your mentors. Pursuing what "you want to do" is somewhat of a childish fantasy. What do you know about what you want to do? You have significantly less experience than your mentors. This applies at any level, even founders have mentors! If, at any stage of your journey, you cannot humble yourself and receive guidance, you will fail, even if you're lucky for a while. If you are regularly soliciting feedback from your mentors, they will be able to point out the areas in which you have aptitude. Note the plural use of _mentors_ - a single opinion is an opinion, multiple opinions in agreement form a consensus.

### Enjoying Something

As stated before coding is a given. You must love coding. One thing isn't enough. What else excites you, design, product, infrastructure, hardware...? Cultivate this "second interest" as you did with coding.

If you go after what you enjoy but you're not good at it even after trying to get good at it, you become sort of a goofy idealist that nobody takes seriously and with no firepower to back up your ideas. If you go after what you're good at but you don't enjoy it you'll end up being resentful. This resentment _will_ come out in innumerable ways through your professional or personal life. It may not undo you, but why would you want that life for yourself?

When it comes time to specialize, management, or IC track, make sure you enjoy what you're going to be doing and that you're good at it. You need both! Got it? Now go forth and make some good mistakes.
