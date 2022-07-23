import React from 'react';
import './App.css';
import { Header } from './components/Header/Header';
import { TodoList } from './components/TodoList/TodoList';

function App() {
  return (
    <div className="App">
      <Header />
      <TodoList todos={[
        {title: "ねる", isCompleted: true},
        {title: "テストする", description: "マイクロソフト", isCompleted: true},
        {title: "起きる", isCompleted: false}
      ]} />
    </div>
  );
}

export default App;
