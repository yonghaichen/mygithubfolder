The rectangular visual mode is very useful in temporarily commenting code, or in removing comments.

Suppose we had a code block like

        int main()
        {
                register int k;
                ...
                ...
        }

Put your cursor at the start of the main declaration. Hit **Ctrl+v** followed by **6j**. Then hit **Shift+i** and type `//`. Finally hit **Esc** to come out of the visual mode. This should comment the entire block with C++ style comments.