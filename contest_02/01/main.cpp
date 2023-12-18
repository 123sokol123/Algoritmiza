bool is_prime(int number) {
    bool flag = true;
    int d=2;
    while (d * d < number) {
        if (number% d == 0){
            return false;
        }
        d++;
    }
    if (number % d == 0) {
        return false;
    }
    
    return true;
}
