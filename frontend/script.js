// HOST ADDRESS
// const ADDR = "http://3.72.23.143"
const ADDR = "http://127.0.0.1:7696"
const HABITS_DIV = document.getElementById("habits")

// PUT HABIT
const updateHabit = async (id, button, habit, habitSpan) => {
    const data = {method: "PUT"}
    const response = await fetch(`${ADDR}/habits/${id}`, data)
    
    if (!response.ok) {
      console.log(`status: ${response.status} text: ${response.statusText}`)
      return
    }

    button.remove()
    habitSpan.innerHTML = `${habit.Name} ${habit.Days+=1}`
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
const createHabit = (habit) => {
    let out = []
    
    const habitSpan = document.createElement("span")
    habitSpan.innerHTML = `${habit.Name} ${habit.Days}`
    out.push(habitSpan)
    
    if (!habit.Inc) {
        const button = document.createElement("button")
        button.innerHTML = "+"
        button.onclick = () => { updateHabit(habit.Id, button, habit, habitSpan) }
        out.push(button)
    }

    const br = document.createElement("br")
    out.push(br)

    return out
}

// ADD HABITS TO DOM
const addToDom = (habits) => {
  for (const habit of habits) {
    const h = createHabit(habit)
    for (const e of h) {
      HABITS_DIV.append(e)
    }
  }
}

// ON STARTUP
const main = async () => {
    const habits = await getHabits()
    addToDom(habits)
}

main()
