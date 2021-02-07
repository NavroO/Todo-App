import React, { useState } from 'react';


const TodoList = () => {

    const [Todo] = useState('Hello from useState');

    <div>
        <strong>{Todo}</strong>
    </div>

}

export default TodoList;