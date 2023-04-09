import { Button, Grid, TextField } from '@mui/material';
import { useEffect, useState } from 'react';
import AddIcon from '@mui/icons-material/Add';
import TodoItem from './TodoItem';
import axios from 'axios';

const TodoList = () => {
    const [todo, setTodo] = useState('');
    const [todos, setTodos] = useState([]);

    const handleChangeTodo = (event) => {
        setTodo(event.target.value);
    }
    const getTodosData = async() => {
        await axios.get("http://localhost:8080/todo/api/v1/todos")
            .then((res) => {
                setTodos(res.data);
            }).catch((err) => {
                console.error(err);
            })
    }

    useEffect(() => {
        getTodosData();
    }, []);

    const handleCreateTodo = async () => {
        const params = new URLSearchParams;
        params.append("title", todo);
        await axios.post("http://localhost:8080/todo/api/v1/todos", params)
            .then(() => {
                getTodosData();
                setTodo('');
            })
            .catch((err) => {
                console.error(err);
            });
        console.log("create todo");
    }

    const handleDeleteTodo = async (id) => {
        await axios.delete(`http://localhost:8080/todo/api/v1/todos/${id}`)
            .then(() => {
                getTodosData();
            })
            .catch((err) => {
                console.error(err);
            });
    }

    return (
        <>
            <Grid container alignItems="center" justifyContent="center" sx={{ m: 3 }}>
                <TextField
                    label="todo"
                    name='todo'
                    value={todo}
                    type="text"
                    variant="outlined"
                    size='small'
                    onChange={handleChangeTodo}
                />
                <Button variant="contained" endIcon={<AddIcon />} onClick={ handleCreateTodo }>CREATE</Button>
            </Grid>
            <TodoItem {...todos} onClick={ handleDeleteTodo } />
        </>
    );
}

export default TodoList;
