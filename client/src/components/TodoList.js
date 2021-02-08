import React, { useState } from 'react';
import './Todo.css'

const TodoList = () => {

    const [Todo, setTodo] = useState([
        "Pies",
        "Kot",
        "Auto"
    ]);

    const addItem = () => {
        const getValueFromInput = document.getElementsByClassName("events__Container__input").value;
        alert(getValueFromInput);
    }

    return (
        
        <div className="eventsConteiner">
            <input type="text" placeholder="Enter new todo" onChange={e => setTodo(e.target.value)}/>
            <button onClick={addItem}>Add new todo</button>
            {Todo.map((Todo) => {
                return (
                    <div className="todo">
                        {Todo}
                        <button className="todo__button">Remove</button>
                    </div>
                );
            })}
        </div>
    );
}

export default TodoList;