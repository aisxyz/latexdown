# latexdown

latexdown 用于解析与数学相关的 LaTex，目前只处理了一部分。

# 备注

这只是 LaTex 处理的一种风格尝试，很多处理过程并没有照顾到所有的 LaTex 命令，这也是为什么在解析过程中在遇到不支持的 LaTex 命令时直接 panic 而非返回 error 的原因，目的就是为了遇到时能在相应的位置加上处理逻辑。但总的来说，骨架就是目前这个样子，添加新的处理节点时只需按照已有的节点仿造即可。
