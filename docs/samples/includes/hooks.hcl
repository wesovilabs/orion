before smoke-test {
    print {
        msg = "hook before: smoke-test"
    }
}
after each {
    print {
        msg = "hook after: each"
    }
}