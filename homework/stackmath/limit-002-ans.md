Using Stolz theorem and definite integral    
Let:    
$$
\begin{align*}
A&=\lim_{n \to \infty}\frac{1+\sqrt{3\cdot4}+...+\sqrt[n]{(n+1)(n+2)...(2n)}}{n^2} \\
A&=\lim_{n \to \infty}\frac{\sqrt[n]{(n+1)(n+2)...(2n)}}{n^2-(n-1)^2}\\
A&=\lim_{n \to \infty}\frac{1}{2}\cdot\sqrt[n]{(1+\frac{1}{n})(1+\frac{2}{n})...(1+\frac{n}{n})}
\end{align*}
$$
Let: 
$$
\begin{align*}
B&=\sqrt[n]{(1+\frac{1}{n})(1+\frac{2}{n})...(1+\frac{n}{n})}\\
ln(B)&=\frac{1}{n}\sum_{i=1}^{n}ln(1+\frac{i}{n})\\
C&=\lim_{n \to \infty}lnB=\int_{0}^{1}ln(1+x)dx\\
C&=(1+x)ln(1+x)|_{0}^{1}-x|_{0}{1}=2ln(2)-1
\end{align*}
$$
Hence:   
$$
\begin{align*}
A&=\frac{2}{e}
\end{align*}
$$



