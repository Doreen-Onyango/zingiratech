document.addEventListener('DOMContentLoaded', function() {
    const auth = firebase.auth();
    
    // Login form handler
    const loginForm = document.querySelector('.auth-form');
    if (loginForm && !loginForm.id) { // Check if it's the login form
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;
            
            try {
                const userCredential = await auth.signInWithEmailAndPassword(email, password);
                const user = userCredential.user;
                
                // Get ID token
                const idToken = await user.getIdToken();
                
                // Send token to your backend
                const response = await fetch('/api/auth/verify', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${idToken}`
                    }
                });
                
                if (response.ok) {
                    // Store the token in localStorage
                    localStorage.setItem('authToken', idToken);
                    // Redirect to dashboard on success
                    window.location.href = '/dashboard';
                } else {
                    throw new Error('Authentication failed');
                }
            } catch (error) {
                console.error('Login error:', error);
                showNotification('error', error.message);
            }
        });
    }

    // Signup form handler
    const signupForm = document.getElementById('signup-form');
    if (signupForm) {
        signupForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            
            const submitButton = signupForm.querySelector('.auth-button');
            const errorMessage = document.getElementById('error-message');
            
            // Get form values
            const fullname = document.getElementById('fullname').value;
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;
            const confirmPassword = document.getElementById('confirm-password').value;
            
            // Basic validation
            if (password !== confirmPassword) {
                errorMessage.textContent = 'Passwords do not match';
                errorMessage.classList.add('show');
                return;
            }
            
            if (password.length < 6) {
                errorMessage.textContent = 'Password must be at least 6 characters';
                errorMessage.classList.add('show');
                return;
            }
            
            try {
                // Show loading state
                submitButton.classList.add('loading');
                submitButton.disabled = true;
                errorMessage.classList.remove('show');
                
                // Create user
                const userCredential = await auth.createUserWithEmailAndPassword(email, password);
                const user = userCredential.user;
                
                // Update profile
                await user.updateProfile({
                    displayName: fullname
                });
                
                // Get ID token
                const idToken = await user.getIdToken();
                
                // Store token
                localStorage.setItem('authToken', idToken);
                
                // Show success message
                showNotification('success', 'Account created successfully!');
                
                // Redirect to dashboard
                setTimeout(() => {
                    window.location.href = '/dashboard';
                }, 1500);
                
            } catch (error) {
                console.error('Signup error:', error);
                errorMessage.textContent = error.message;
                errorMessage.classList.add('show');
                submitButton.classList.remove('loading');
                submitButton.disabled = false;
            }
        });
    }
    
    // Password visibility toggle
    const toggleButtons = document.querySelectorAll('.toggle-password');
    toggleButtons.forEach(button => {
        button.addEventListener('click', function() {
            const input = this.parentElement.querySelector('input');
            const type = input.getAttribute('type') === 'password' ? 'text' : 'password';
            input.setAttribute('type', type);
            
            const icon = this.querySelector('i');
            icon.classList.toggle('fa-eye');
            icon.classList.toggle('fa-eye-slash');
        });
    });
    
    // Notification system
    function showNotification(type, message) {
        const notification = document.createElement('div');
        notification.className = `notification ${type}`;
        notification.innerHTML = `
            <i class="fas ${type === 'success' ? 'fa-check-circle' : 'fa-exclamation-circle'}"></i>
            <p>${message}</p>
        `;
        
        document.body.appendChild(notification);
        
        setTimeout(() => {
            notification.remove();
        }, 3000);
    }
}); 