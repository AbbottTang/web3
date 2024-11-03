// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

contract Todos {
    struct Todo {
        string text;
        bool completed;
    }
    //an array of 'Todo' structs
    Todo[] public todos;

    function create(string calldata _text) public {
        //初始化struct的三种方式
        //像调用函数一样调用它
        todos.push(Todo(_text, false));

        //键值映射
        todos.push(Todo({text: _text, completed: false}));

        //初始化一个空的然后更新他
        Todo memory todo;
        todo.text = _text;
        //bool 默认值为false
        todos.push(todo);
    }

    // Solidity自动为todos创建了一个getter
    //实际上你并不需要这个函数。
    function get(
        uint256 _index
    ) public view returns (string memory text, bool completed) {
        Todo storage todo = todos[_index];
        return (todo.text, todo.completed);
    }

    //update text
    function updateText(uint256 _index, string calldata _text) public {
        Todo storage todo = todos[_index];
        todo.text = _text;
    }

    //update completed
    function toggleCompleted(uint256 _index) public {
        Todo storage todo = todos[_index];
        todo.completed = !todo.completed;
    }
}
