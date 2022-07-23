import { Todo } from "../../entities/Todo";
import { TodoItem } from "../TodoItem/TodoItem";
import "./TodoList.scss"

type Props = {
  todos: Todo[];
}

export const TodoList: React.FC<Props> = ({ todos }) => {
  const listItems = todos.map((todo) =>
    <li>
        <TodoItem
          title={todo.title}
          description={todo.description}
          isCompleted={todo.isCompleted}
        />
    </li>
  );

  return (
      <ul className="todo-list">
          {listItems}
      </ul>
  )
}