import Foundation

let a = DispatchSemaphore(value: 0)
DispatchQueue.global(qos: .background).async {
    a.signal()
    print("signalled")
}

a.wait()
print("exiting")