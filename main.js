(async() => {
    function wait(ms) {
        return new Promise(resolve => setTimeout(() =>resolve(), ms));
    };

    let i = 0

    while(true) {
        const axios = require('axios').default;
        const id = "its your toss id"
        await axios.get(`https://toss.me/${id}`, {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36'
        })
        i++

        console.log(i)
    }
})()
