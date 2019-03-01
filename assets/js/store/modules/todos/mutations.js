
export default {
    addTodo (state, todo) {
        state.items.push(todo)
    },

    removeTodo (state, todo) {
        state.items.splice(state.items.indexOf(todo), 1)
    },

    editTodo (state, { todo, text = todo.text, done = todo.done }) {
        todo.text = text
        todo.done = done
    }
};