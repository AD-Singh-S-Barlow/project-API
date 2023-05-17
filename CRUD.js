const url = //Wherever your database is
options = {
    method : "DELETE"
};

//Deleting an item from the database
fetch(url, options)
    .then(response => {
        if (!response.ok) {
            throw new Error("No Response from Network");
        }
    })
    .catch(error => {
        console.error("Error deleting item:", error);
    });


//Adding an item to the database

function addData(age, firstName, lastName){
    const data = {age, firstName, lastName}
    //Using react here, subject to change
    axios.post("LOCATION OF DATABASE GOES HERE", data)
        .then(response => {
            console.log(response.data);
        })
        .catch(error => {
            console.error(error);
        });
}

//Updating the Database
