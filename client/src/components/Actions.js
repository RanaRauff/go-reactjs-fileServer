import React from "react"
import axios from "axios"
import {alert} from "react-alert"


let endpoint="http://0.0.0.0:8081"

class Actions extends React.Component{
    constructor(props){
        super(props);

        this.state={
            file:null,
            filedata:[]
        }
    }

    componentDidMount(){
        console.log("THE COMPONENT DID MOUNTED")
    }

    handleUploadChange=(e)=>{
        e.preventDefault();
        console.log("this is the event inside fun handleUpload",e.target,e.target,e.target.files)
        this.setState({file:e.target.files[0]})
        console.log("This is the updated DB",this.state)
        
        
    }
    handleUploadSubmit=(e)=>{
        e.preventDefault()
        const response = new FormData()
        response.append("files",this.state.file,this.state.file.name)
        console.log("This is the handleUploadSubmit",this.state,response)
        axios.post(endpoint+"",response).then((res)=>{
            if (res.status==200){
                this.setState({file:null})
                
                window.alert("file uploaded successfully...")
                window.location.reload()
            }else{
                window.alert("Error in Uploading File...")
                console.log(res.status)
            }
        })
    }
    handleReadData=()=>{
        axios.get(endpoint,"/",).then(res=>this.setState({filedata:res.data}))
        console.log("The Data is ", this.state)
    }

    handleDownload=(ele)=>{
        let response = {
            data:ele.target.value
        }
        console.log("This is the download data", ele.target.value , response)
        axios.post(endpoint+"/dwnld",{data:ele.target.value}).then(console.log("Data Sent"))
    }

    render(){
        console.log("this is the state inside render",this.state)
        return(

            <div>
                {/* kdnfsdnfksdn */}
            <form onSubmit={this.handleUploadSubmit}> 
            
                <input type="file" name="files" placeholder="files" onChange = {this.handleUploadChange} />
                <input type="submit"></input>
            </form>

            <button onClick={this.handleReadData}>Read Data</button>
            <br/>
            {this.state.filedata.map(

                (obj)=>
                <div key = {obj._id}>
                    <img height="200"  width="400" src={endpoint+"/statics/"+obj.name}></img>
                    <button onClick={this.handleDownload} value={obj.name}>Download</button>
                    <br/>
                </div>    
            
            )}
            </div>


        )
    }
}

export default Actions;