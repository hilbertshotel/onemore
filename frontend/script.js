// HOST ADDRESS
// const ADDR = "http://3.72.23.143"
const ADDR = "http://127.0.0.1:7696"
const HABITS_DIV = document.getElementById("habits")

// PUT HABIT
const updateHabit = async (habit, habitDiv) => {
    const response = await fetch(`${ADDR}/habits/${habit.Id}`, {method: "PUT"})
    if (!response.ok) {
      console.log(`status: ${response.status} text: ${response.statusText}`)
      return
    }

    habitDiv.innerHTML = `${habit.Name} = ${habit.Days+=1}`
    habitDiv.id = ""
}

// POST HABIT
const postHabit = async () => {
    const inputValue = document.getElementById("new_habit").value
    if (inputValue === "") {
      return
    }

    const data = { 
      method: "POST",
      headers: { 
        "Content-Type": "application/json"
      },
      body: JSON.stringify(inputValue)
    }

    const response = await fetch(`${ADDR}/habits/`, data)
    if (!response.ok) {
      console.log(`status: ${response.status} text: ${response.statusText}`)
      return
    }

    const habit = await response.json()
    addToDom([habit])
}

// GET HABIT
const getHabits = async () => {
    const response = await fetch(`${ADDR}/habits`)
    if (!response.ok) {
        console.log(`status: ${response.status} text: ${response.statusText}`)
        return []
    }
    const habits = await response.json()
    return habits
}

// CREATE HABIT
const create = (habit) => {
    const habitDiv = document.createElement("div")
    habitDiv.innerHTML = `${habit.Name} = ${habit.Days}`
    habitDiv.className = "habit"

    if (!habit.Inc) {
      habitDiv.id = "no_inc"
      habitDiv.onclick = () => { updateHabit(habit, habitDiv) }
    }
    
    return habitDiv
}

// ADD HABITS TO DOM
const addToDom = (habits) => {
  for (const habit of habits) {
    HABITS_DIV.append(create(habit))
  }
}

// ON STARTUP
const main = async () => {
    const habits = await getHabits()
    addToDom(habits)
}

main()
