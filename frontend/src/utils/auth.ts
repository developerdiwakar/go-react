import { jwtDecode } from 'jwt-decode'
// Set auth data for User including
export const setAuthToken = (token: string) => {
    localStorage.setItem('auth-token', token);
}

// Get auth token from Local storage
export const getToken = ():string => {
    const tokenStr = localStorage.getItem('auth-token')
    return !tokenStr ? "" : tokenStr
}


// Remove the token and user from the Local storage
export const removeAuthToken = () => {
    localStorage.removeItem('auth-token');
}

// Check the token expiry time
export const checkIsTokenExpired = ():boolean => {
    // Extract the token
    if(getToken() == ""){
        return true
    }
    const decodedToken = jwtDecode(getToken()) as { [key: string]: any };
    return decodedToken.exp < Date.now()/1000;
}

// Check if the user login session expired or not
export const checkIsLoggedIn = () => {
    return ! checkIsTokenExpired()
}
