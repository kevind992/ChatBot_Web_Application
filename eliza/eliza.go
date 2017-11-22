package eliza

import(

	"regexp"
	"time"
	"math/rand"
	//"fmt"
	"strings"

)


func ElizaStart(userInput string) string{

	rand.Seed(time.Now().UTC().UnixNano())

	//userInput = Reflect(userInput)

	ranNum := rand.Intn(3)

	re1 := regexp.MustCompile(`(?im)^\s*Hi, my name is ([^.,!?]*)[.,!?]?`)
	if re1.MatchString(userInput){
		return re1.ReplaceAllString(userInput,"Hi $1, how are you today?")
	}
	re1 = regexp.MustCompile(`(?im)^\s*I need ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a1 := []string{
			`Why do you need $1?`,
			`Would it really help you to get $1?`,
			`Are you sure you need $1?`,}

		return re1.ReplaceAllString(userInput, a1[ranNum])
	}

	re1 = regexp.MustCompile(`(?im)^\s*Why don'?t you ([^\?]*)\?*\s*$`)
	if re1.MatchString(userInput){
		a2 := []string{
			"Do you really think I don't $1?",
			"Perhaps eventually I will $1?",
			"Do you really want me to $1?",}
			return re1.ReplaceAllString(userInput, a2[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*Why can'?t I ([^\?]*)\?*\s*$`)
	if re1.MatchString(userInput){
		a3 := []string{
			"Do you think you should be able to $1?",
			"If you could $1, what would you do?",
			"I don't know - why can't you $1?",
			"Have you really tried?",}
			return re1.ReplaceAllString(userInput, a3[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*I can'?t ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a4 := []string{
			"How do you know you can't $1",
			"Perhaps you could $1 if you tried.",
			"What would it take for you to $1",}
			return re1.ReplaceAllString(userInput, a4[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*I am ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a5 := []string{
			"Did you come to me because you are $1",
			"How long have you been $1",
			"How do you feel about being $1",}
			return re1.ReplaceAllString(userInput, a5[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*I'm ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a6 := []string{
			"How does being $1 make you feel",
			"Do you enjoy being $1",
			"Why do you tell me you're $1",
			"Why do you think you're $1",}
			return re1.ReplaceAllString(userInput, a6[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*How [^\?]*\?*\s*$`)
	if re1.MatchString(userInput){
		a7 := []string{
			"How do you suppose?",
			"Perhaps you can answer your own question.",
			"What is it you're really asking",}
			return re1.ReplaceAllString(userInput, a7[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*Yes\s*$`)
	if re1.MatchString(userInput){
		a8 := []string{
			"You seem quite sure.",
			"Interesting...",
			"OK, but can you elaborate a bit?",}
			return re1.ReplaceAllString(userInput, a8[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*I have ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a9 := []string{
			"Why do you tell me you've $1",
			"Have you really $1",
			"Now that you have $1, what will you do next?",}
			return re1.ReplaceAllString(userInput, a9[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*Is there ([^\?]*)\?*\s*$`)
	if re1.MatchString(userInput){
		a10 := []string{
			"Do you think there is $1",
			"It's likely that there is $1",
			"Would you like there to be",}
			return re1.ReplaceAllString(userInput, a10[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*You ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a11 := []string{
			"We should be discussing you, not me.",
			"Why do you say that about me?",
			"Why do you care whether I $1?",}
			return re1.ReplaceAllString(userInput, a11[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*Why ([^\?]*)\?*\s*$`)
	if re1.MatchString(userInput){
		a12 := []string{
			"Why don't you tell me the reason why $1?",
			"Why do you think $1?",}
			return re1.ReplaceAllString(userInput, a12[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*I want ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a13 := []string{
			"What would it mean to you if you got $1",
			"Why do you want $1",
			"What would you do if you got $1?",
			"If you got $1, then what would you do?",}
			return re1.ReplaceAllString(userInput, a13[ranNum])
		
	}
	re1 = regexp.MustCompile(`[^\?]*\?*\s*$`)
	if re1.MatchString(userInput){
		a14 := []string{
			"Why do you ask that?",
			"Please consider wheather you can answer your own question.",
			"Perhaps the answer lies within yourself?",
			"Why don't you tell me?",}
			return re1.ReplaceAllString(userInput, a14[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*quit\s*$`)
	if re1.MatchString(userInput){
		a15 := []string{
			"Thank you for visiting.",
			"Bye now, thanks for visiting.",
			"Have a nice day.",}
			return re1.ReplaceAllString(userInput, a15[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*You are ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a16 := []string{
			"Why do you think I am $1?",
			"Does it please you to think that I'm $1",
			"Perhaps you would like me to be $1",
			"Perhaps you're really talking about yourself?",}
			return re1.ReplaceAllString(userInput, a16[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*You're ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a17 := []string{
			"Why do you say I am $1?",
			"Why do you think I am $1?",
			"Are we talking about you, or me?",}
			return re1.ReplaceAllString(userInput, a17[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*I don't ([^\.!]*)[\.!]*\s*$`)
	if re1.MatchString(userInput){
		a18 := []string{
			"Don't you really $1?",
			"Why don't you $1",
			"Do you want to $1?",}
			return re1.ReplaceAllString(userInput, a18[ranNum])
		
	}
	re1 = regexp.MustCompile(`(?im)^\s*quit\s*$`)
	if re1.MatchString(userInput){
		a19 := []string{
			"Thank you for visiting.",
			"Bye now, thanks for visiting.",
			"Have a nice day.",}
			return re1.ReplaceAllString(userInput, a19[ranNum])
		
	}
	re1 = regexp.MustCompile(`(.*)`)
	if re1.MatchString(userInput){
		a20 := []string{
			"Please tell me more.",
			"Let's change focus a bit... Tell me about your family",
			"Can you elaborate on that?",
			"I see.",
			"Why do you that $1?",
			"Very interesting.",
			"$1",
			"I see, and what does that tell you?",
		}
			return re1.ReplaceAllString(userInput, a20[ranNum])
		
	}
	return ""
}
func Reflect(input string)string{

	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	reflections := [][]string{
		{`am`,`are`},
		{`was`,`were`},
		{`i`,`you`},
		{`i'd`,`you would`},
		{`i've`,`you have`},
		{`i'll`,`you will`},
		{`my`,`your`},
		{`are`,`am`},
		{`you've`,`I have`},
		{`you'll`,`I will`},
		{`your`,`my`},
		{`yours`,`mine`},
		{`you`,`me`},
		{`me`,`you`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}
