/*
#include <iostream>
#include <vector>
#include <iomanip>
#include <algorithm>

using namespace std;

class Bank {
    public:
        virtual void assignLoans(vector<int> loans) = 0;
        virtual void averageLoan() = 0;
        virtual void minLoan() = 0;
        virtual void maxLoan() = 0;
};

class AbsBankLoans

class PersonalLoanDept : public Bank {
    private:
        vector<int> loanAmounts;

    public:
        PersonalLoanDept(int clients) {
            loanAmounts.resize(clients, 0);
        }

        void assignLoans(vector<int> loans) override {
            int size = min((int)loanAmounts.size(), (int)loans.size());
            for (int i = 0; i < size; i++) {
                loanAmounts[i] = loans[i];
            }
            cout << "Loans for clients processed" << endl;
        }

        void averageLoan() override {
            double sum = 0;
            int count = 0;
            for (int amt : loanAmounts) {
                sum += amt;
                count++;
            }
            cout << fixed << setprecision(2);
            cout << "Average loan Amount for clients is " << (count ? sum / count : 0) << endl;
        }

        void maxLoan() override {
            int maxAmt = *max_element(loanAmounts.begin(), loanAmounts.end());
            cout << "Minimum Loan amount amongst clients is " << maxAmt << endl;
        }

        void minLoan() override {
            int minAmt = *min_element(loanAmounts.begin(), loanAmounts.end());
            cout << "Minimum Loan amount amongst clients is " << minAmt << endl;
        }
};

class BusinessLoanDept : public Bank {
    private:
        vector<int> loanAmounts;

    public:
        BusinessLoanDept(int businesses) {
            loanAmounts.resize(businesses, 0);
        }

        void assignLoans(vector<int> loans) override {
            int size = min((int)loanAmounts.size(), (int)loans.size());

            for (int i = 0; i < size; i++) {
                loanAmounts[i] = loans[i];
            }
            cout << "Loans for business processed" << endl;
        }

        void averageLoan() override {
            double sum = 0;
            int count = 0;
            for (int amt : loanAmounts) {
                sum += amt;
                count++;
            }
            cout << fixed << setprecision(2);
            cout << "Average loan Amount for clients is " << (count ? sum / count : 0) << endl;
        }

        void maxLoan() override {
            int maxAmt = *max_element(loanAmounts.begin(), loanAmounts.end());
            cout << "Minimum Loan amount amongst clients is " << maxAmt << endl;
        }

        void minLoan() override {
            int minAmt = *min_element(loanAmounts.begin(), loanAmounts.end());
            cout << "Minimum Loan amount amongst clients is " << minAmt << endl;
        }
};

int main() {
    int n, m;

    cin >> n >> m;

    vector<int> personalLoans(n);
    vector<int> businessLoans(m);

    for (int i = 0; i < n; ++i) {
        cin >> personalLoans[i];
    }

    for (int i = 0; i < m; ++i) {
        cin >> businessLoans[i];
    }

    PersonalLoanDept p(n);
    BusinessLoanDept b(m);

    p.assignLoans(personalLoans);
    b.assignLoans(businessLoans);

    p.averageLoan();
    p.maxLoan();
    p.minLoan();

    b.averageLoan();
    b.maxLoan();
    b.minLoan();

    return 0;
}
*/

#include <iostream>
#include <vector>
#include <iomanip>
#include <algorithm>

using namespace std;

class Bank {
public:
    virtual void assignLoans(vector<int> loans) = 0;
    virtual void averageLoan() = 0;
    virtual void minLoan() = 0;
    virtual void maxLoan() = 0;
};

class AbsBankLoans : public Bank {
    protected:
        vector<int> loanAmounts;

    public:
        void averageLoan() override {
            double sum = 0;
            int count = 0;
            for (int amt : loanAmounts) {
                sum += amt;
                count++;
            }
            cout << fixed << setprecision(2);
            cout << "Average loan Amount is " << (count ? sum / count : 0) << endl;
        }
};

class PersonalLoanDept : public AbsBankLoans {
    public:
        PersonalLoanDept(int clients) {
            loanAmounts.resize(clients, 0);
        }

        void assignLoans(vector<int> loans) override {
            int size = min((int)loanAmounts.size(), (int)loans.size());
            for (int i = 0; i < size; i++) {
                loanAmounts[i] = loans[i];
            }
            cout << "Loans for personal department processed." << endl;
        }

        void minLoan() override {
            int minAmt = *min_element(loanAmounts.begin(), loanAmounts.end());
            cout << "Minimum Loan amount is " << minAmt << endl;
        }

        void maxLoan() override {
            int maxAmt = *max_element(loanAmounts.begin(), loanAmounts.end());
            cout << "Maximum Loan amount is " << maxAmt << endl;
        }
};

class BusinessLoanDept : public AbsBankLoans {
    public:
        BusinessLoanDept(int businesses) {
            loanAmounts.resize(businesses, 0);
        }

        void assignLoans(vector<int> loans) override {
            int size = min((int)loanAmounts.size(), (int)loans.size());
            for (int i = 0; i < size; i++) {
                loanAmounts[i] = loans[i];
            }
            cout << "Loans for business department processed." << endl;
        }

        void minLoan() override {
            int minAmt = *min_element(loanAmounts.begin(), loanAmounts.end());
            cout << "Minimum Loan amount is " << minAmt << endl;
        }

        void maxLoan() override {
            int maxAmt = *max_element(loanAmounts.begin(), loanAmounts.end());
            cout << "Maximum Loan amount is " << maxAmt << endl;
        }
};

int main() {
    int n, m;

    cin >> n >> m;

    vector<int> personalLoans(n);
    vector<int> businessLoans(m);

    for (int i = 0; i < n; ++i) {
        cin >> personalLoans[i];
    }

    for (int i = 0; i < m; ++i) {
        cin >> businessLoans[i];
    }

    PersonalLoanDept p(n);
    BusinessLoanDept b(m);

    p.assignLoans(personalLoans);
    b.assignLoans(businessLoans);

    p.averageLoan();
    p.maxLoan();
    p.minLoan();

    b.averageLoan();
    b.maxLoan();
    b.minLoan();

    return 0;
}

/*

#include <iostream>
#include <vector>
#include <regex>
#include <string>
#include <unordered_map>
#include <deque>

using namespace std;

class IPValidator {
    private:
        vector<regex> blacklistPatterns;
        unordered_map<string, deque<int>> ipHistory;

        string wildcardToRegex(const string& pattern) {
            string regexPattern;
            for (char ch : pattern) {
                if (ch == '*') regexPattern += ".*";
                else if (ch == '.') regexPattern += "\\.";
                else regexPattern += ch;
            }
            return regexPattern;
        }

        bool isBlacklisted(const string& ip) {
            for (const auto& r : blacklistPatterns) {
                if (regex_match(ip, r)) {
                    return true;
                }
            }
            return false;
        }

        bool exceedsRateLimit(const string& ip, int currentTime) {
            auto& times = ipHistory[ip];

            while (!times.empty() && times.front() == currentTime - 5) {
                times.pop_front();
            }

            return times.size() >= 2;
        }

    public:
        IPValidator(const vector<string>& blacklisted) {
            for (const auto& pat : blacklisted) {
                blacklistPatterns.push_back(regex(wildcardToRegex(pat)));
            }
        }

        vector<int> validate(const vector<string>& requests) {
            vector<int> result;

            for (int t = 0; t < requests.size(); t++) {
                const string& ip = requests[t];
                bool blocked = false;

                if (isBlacklisted(ip)) {
                    blocked = true;
                } else if (exceedsRateLimit(ip, t)) {
                    blocked = true;
                }

                result.push_back(blocked ? 1 : 0);

                if (!blocked) {
                    ipHistory[ip].push_back(t);
                }
            }

            return result;
        }
};

int main() {
    vector<string> blacklist = {"111.*.255", "12.*"};
    vector<string> requests = {"121.3.5.255", "12.13.5.255", "111.3.5.255", "121.3.5.255"};

    IPValidator validator(blacklist);
    vector<int> result = validator.validate(requests);

    for (int val : result) {
        cout << val << endl;
    }

    return 0;
}
*/