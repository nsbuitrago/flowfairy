
export const Ping = async () => {
    const route = "http://localhost:8000/status";
    const response = await fetch(route, {
        method: 'GET'
    })
}
export const LoadFlow = async (fileList) => {
    console.log(fileList);

    const route = "http://localhost:8000/load";
    let requestBody = new FormData();

    for (let i=0; i < fileList.length; i++) {
        requestBody.append(`fcs_file_${i}`, fileList[i])
    }

    const upload = await fetch(route, {
        method: 'POST',
        body: requestBody
    })

    console.log(upload.json());
}