// server address
const ADDR = "http://3.72.23.143"

// update habit
const updateHabit = (id, button, habit, habitSpan) => {
    
     

    button.remove()
    habitSpan.innerHTML = `${habit.Name} ${habit.Days+=1}`
    console.log(id)
}

// post habit
const postHabit = () => {
    console.log("post")
}

// get habits
const getHabits = async () => {
    const response = await fetch(`${ADDR}/habits`)
    if (response.ok) {
        const habits = await response.json()
        return habits
    }
    console.log(`status: ${response.status} text: ${response.statusText}`)
}

// on page start
const main = async () => {
    const habitsDiv = document.getElementById("habits")
    const habits = await getHabits()
    
    for (const habit of habits) {
        const habitSpan = document.createElement("span")
        const br = document.createElement("br")
        const button = document.createElement("button")
        habitSpan.innerHTML = `${habit.Name} ${habit.Days}`
        
        habitsDiv.append(habitSpan)
        if (!habit.Inc) {
            const button = document.createElement("button")
            button.innerHTML = "+"
            button.onclick = () => { updateHabit(habit.Id, button, habit, habitSpan) }
            habitsDiv.append(button)
        }
        habitsDiv.append(br)
    }

}

main()

