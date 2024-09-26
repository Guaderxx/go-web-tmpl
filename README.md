# gowebtmpl

Golang web template.

A Go(Golang) Web Backend Template project with [gin][gin], [ent][ent], [jwt][jwt], [sonyflake][sonyflake], [cobra][cobra], [viper][viper].


## Architecture of the project

I was trying to design a clean project like you can esaily replacement of components.  

The truth is, I can't.

Still, you can replace some componet like `gin`, `logger(slog)`, `ent`, `sonyflake` .    
Just need some time to fix the interface signature and function signature.



[gin]: https://github.com/gin-gonic/gin
[ent]: https://github.com/ent/ent
[jwt]: https://github.com/golang-jwt/jwt
[cobra]: https://github.com/spf13/cobra
[viper]: https://github.com/spf13/viper
[sonyflake]: https://github.com/sony/sonyflake
[gonew]: https://pkg.go.dev/golang.org/x/tools/cmd/gonew

