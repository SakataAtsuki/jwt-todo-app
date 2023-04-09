import { Button, Grid, List, ListItem, ListItemText } from "@mui/material"
import { Box } from "@mui/system";
import WorkIcon from '@mui/icons-material/Work';
import DeleteIcon from '@mui/icons-material/Delete';

const TodoItem = (props) => {
    const todos = props.todos;

    return (
        <Grid container alignItems="center" justifyContent="center">
            <Box
                sx={{ width: '100%', height: 400, maxWidth: 360, bgcolor: 'background.paper' }}>
                <List
                    sx={{
                        width: '100%',
                        maxWidth: 360,
                        bgcolor: 'background.paper',
                        position: 'relative',
                        overflow: 'auto',
                        maxHeight: 300,
                        '& ul': { padding: 0 },
                    }}>
                    {todos?.map((todo, index) => {
                        return (
                            <ListItem key={index}>
                                <WorkIcon sx={{ p: 2 }} />
                                <ListItemText primary={todo.title} />
                                <Button onClick={() => props.onClick(todo.id)} endIcon={<DeleteIcon />}>DELETE</Button>
                            </ListItem>
                        )
                    })}
                </List>
            </Box>
        </Grid>
    );
}

export default TodoItem;
