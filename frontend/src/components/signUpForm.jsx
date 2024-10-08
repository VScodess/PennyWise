import React, { useState } from "react";

const SignUpForm = ({ closeForm }) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [email, setEmail] = useState("");

    const handleSubmit = async(e) => {
        e.preventDefault();
        try {
            const response = await fetch("http://localhost:8080/api/signup", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username, email, password }),
            });

            if (response.ok) {
                const data = await response.json();
                console.log("User created successfully: ", data);
                closeForm();
            } else {
                console.log("Signup failed");
                alert("Signup failed");
            }
        } catch (error) {
            console.error("Error during Signup:", error);
            alert("An error occurred during Signup. Please try again.");
        }
    };

    return (
        <div className="loginFormBackground" onClick={closeForm}>
            <div className="loginFormContainer" onClick={(e) => e.stopPropagation()}>
                <h2 className="formTitle">Sign Up</h2>
                <form onSubmit={handleSubmit}>
                    <div className="inputGroup">
                        <label htmlFor="username">Username:</label>
                        <input
                            type="text"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                        />
                    </div>
                    <div className="inputGroup">
                        <label htmlFor="password">Password:</label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>
                    <div className="inputGroup">
                        <label htmlFor="email">Email:</label>
                        <input
                            type="email"
                            id="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                        />
                    </div>
                    <button type="submit" className="loginFormSubmitButton">Sign Up</button>
                </form>
            </div>
        </div>
    );
};

export default SignUpForm;
