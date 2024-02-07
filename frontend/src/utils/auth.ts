// Set session data for User including token
export const setAuthToken = (token: string) => {
    sessionStorage.setItem('auth-token', token);
}

// Get auth token from session storage
export const getToken = ():string => {
    const tokenStr = sessionStorage.getItem('auth-token')
    return !tokenStr ? "" : tokenStr
}


// Remove the token and user from the session
export const removeAuthToken = () => {
    sessionStorage.removeItem('auth-token');
}

