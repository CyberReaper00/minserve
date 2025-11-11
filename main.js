let value = 0
const MAX_HUE = 360

const something = setInterval(() => {
	value += 20
	document.body.style.background = `hsl(${value%MAX_HUE}, 50%, 50%)`
}, 400)
