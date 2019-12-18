# What's race condition?

競合状態は2つ以上の操作が正しい順番で実行されなければいけないところで、プログラムが順序を保証するように書かれていなかったときに発生します。
A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.

競合状態はいわゆるデータ競合です。ある並列処理の操作が変数を読み込もうしているときに、ほかの並列処理の操作が不確定のタイミングで同じ変数に書き込みを行おうとすると発生します。
Most of the time, A race condition shows up in what's called a data race, where one concurrent operation attempts to read a variable while at some undermined time another concurrent operation is attempting to write to the same variable.
