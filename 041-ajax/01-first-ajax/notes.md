Looks like the “Make a request” js action in the index.html (in Section 041_AJAX/01) doesn’t execute the alert function (I just opened the index.html file from my file system). The following exception occurs in the developer console:


“Cross-Origin Request Blocked: The Same Origin Policy disallows reading the remote resource at file:///Users/matthew/Documents/go/src/github.com/imattf/gowebdev/041-ajax/01/data.txt. (Reason: CORS request not http).” 


Is this because the file:/// protocol is not supported for XMLHttpRequest?


Work around: I created a simple server that runs from the same directory as the index.html file and now the index.html appears to work as expected w/ $go run main.go (avail here: https://play.golang.org/p/PNyM8kk_69G)

            