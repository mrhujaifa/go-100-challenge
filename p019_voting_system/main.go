package main

import (
	"errors"
	"fmt"
)

/*
===========================================
  Go 100 Challenge
  Problem: #019
  Level: 🟡 Medium
===========================================

Problem: Voting System

Topic:
- struct
- map
- functions
- error handling
- duplicate prevention

Industry Use:
Election Systems / Poll Apps /
Survey Platforms / Decision Making Tools

Rules (English):
- Manage a voting system with candidates
- Each voter can only vote once
- Each candidate has: Name, Party, Votes
- User can:
    * Register as voter (by name + ID)
    * Cast vote for a candidate
    * View live results
    * View winner
    * Exit
- Already voted         => "You have already voted!"
- Candidate not found   => "Candidate not found!"
- Voter not registered  => "Voter not registered!"
- No votes cast yet     => "No votes cast yet!"

Rules (বাংলায়):
- Candidates সহ voting system manage করবে
- প্রতিটা voter শুধু একবারই vote দিতে পারবে
- প্রতিটা candidate এ থাকবে: Name, Party, Votes
- User করতে পারবে:
    * Voter register করা (name + ID দিয়ে)
    * Candidate কে vote দেওয়া
    * Live results দেখা
    * Winner দেখা
    * Exit করা
- Already voted          => "You have already voted!"
- Candidate না পেলে      => "Candidate not found!"
- Voter registered না    => "Voter not registered!"
- কোনো vote নেই        => "No votes cast yet!"

Example Run:
  === Voting System ===
  Candidates:
  1. Rakib  — Party: Blue
  2. Sumaiya — Party: Green
  3. Tanvir — Party: Red

  1. Register Voter
  2. Cast Vote
  3. View Results
  4. View Winner
  5. Exit

  Choose: 1
  Voter Name: Hasan
  Voter ID  : V001
  ✅ Voter registered!

  Choose: 2
  Voter ID  : V001
  Candidate : Rakib
  ✅ Vote cast successfully!

  Choose: 2
  Voter ID  : V001
  ❌ You have already voted!

  Choose: 3
  ================================
  LIVE RESULTS
  ================================
  Rakib    | Blue   | 1 votes
  Sumaiya  | Green  | 0 votes
  Tanvir   | Red    | 0 votes
  ================================

  Choose: 4
  🏆 Winner: Rakib (Blue) with 1 votes!
===========================================
*/

type Candidate struct {
	Name  string
	Party string
	vote  int
}

type Voter struct {
	ID       string
	Name     string
	HasVoted bool
}

type VotingSystem struct {
	candidates map[string]*Candidate
	voters     map[string]*Voter
}

func NewVotingSystem() *VotingSystem {
	return &VotingSystem{
		candidates: map[string]*Candidate{
			"Rakib":   {Name: "Rakib", Party: "Blue"},
			"Sumaiya": {Name: "Sumaiya", Party: "Green"},
			"Tanvir":  {Name: "Tanvir", Party: "Red"},
		},
		voters: make(map[string]*Voter),
	}
}

func (r *VotingSystem) RegisterVoter(voter Voter) error {
	if voter.Name == "" {
		return errors.New("Voter Name cannot be empty!")
	}

	if voter.ID == "" {
		return errors.New("Voter ID cannot be empty!")
	}

	r.voters[voter.ID] = &voter

	return nil
}

func RegisterVoter(r *VotingSystem) {

	var name, Id string

	fmt.Println("Enter your Name:")
	fmt.Scan(&name)
	fmt.Println("Enter your ID (V001):")
	fmt.Scan(&Id)

	err := r.RegisterVoter(Voter{
		ID:   Id,
		Name: name,
	})

	if err != nil {
		fmt.Println("error from register voter:", err)
	}

	fmt.Println("✅ Voter registered!")
}

func (vs *VotingSystem) CastVote(voterId string, candidateName string) error {

	// look up voter by ID
	voter, exists := vs.voters[voterId]
	if !exists {
		return errors.New("Voter not registered!")
	}

	if voter.HasVoted {
		return errors.New("You have already voted!")
	}

	if voter.ID == "" {
		return errors.New("ID cannot be empty!")
	}

	if voter.Name == "" {
		return errors.New("Name cannot be empty!")
	}

	// Check candidate
	candidate, exists := vs.candidates[candidateName]
	if !exists {
		return errors.New("candidate not found")
	}

	// Increase candidate vote
	candidate.vote++

	// mark as voted (actual candidate handling done elsewhere)
	voter.HasVoted = true

	return nil
}

func CastVote(vs *VotingSystem) {

	var voterId, candidateName string
	fmt.Println("Enter your voter ID:")
	fmt.Scan(&voterId)
	fmt.Println("Enter Candidate Name:")
	fmt.Scan(&candidateName)

	err := vs.CastVote(voterId, candidateName)
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	fmt.Println("✅ Vote cast successfully!")
}

func (vs *VotingSystem) ShowResults() {

	fmt.Println("================================")
	fmt.Println("LIVE RESULTS")
	fmt.Println("================================")

	for _, candidate := range vs.candidates {
		fmt.Printf("%-9s | %-6s | %d votes\n",
			candidate.Name,
			candidate.Party,
			candidate.vote,
		)
	}

	fmt.Println("================================")
}

func ShowResults(vs *VotingSystem) {
	vs.ShowResults()
}

//   Choose: 4
//   🏆 Winner: Rakib (Blue) with 1 votes!

func ViewWinner(vs *VotingSystem) {

	max := -1
	var winner *Candidate

	for _, v := range vs.candidates {
		if v.vote > max {
			max = v.vote
			winner = v
		}
	}

	fmt.Printf("🏆 Winner: %s (%s) with %d votes!\n",
		winner.Name,
		winner.Party,
		winner.vote,
	)
}

func mainMenu() {

	votingSystem := NewVotingSystem()

	var choice int
	for {
		fmt.Println("\n=== Voting System ===")
		fmt.Println("Candidates:")
		fmt.Println(" Rakib   — Party: Blue")
		fmt.Println(" Sumaiya — Party: Green")
		fmt.Println(" Tanvir  — Party: Red")
		fmt.Println("-----------------------")
		fmt.Println("1. Register Voter")
		fmt.Println("2. Cast Vote")
		fmt.Println("3. View Results")
		fmt.Println("4. View Winner")
		fmt.Println("5. Exit")
		fmt.Println("-----------------------")
		fmt.Print("Enter your choice (1-5): ")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			RegisterVoter(votingSystem)
		case 2:
			CastVote(votingSystem)
		case 3:
			ShowResults(votingSystem)
		case 4:
			ViewWinner(votingSystem)
		case 5:
			return
		}

	}

}

func main() {
	mainMenu()

}
